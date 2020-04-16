package web

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"../shared"

	"github.com/go-chi/chi"

	"logan/types"

	"../mongo"

	"gopkg.in/mgo.v2/bson"
)

func getLogs(w http.ResponseWriter, r *http.Request) {
	level := chi.URLParam(r, "level")
	if level == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   true,
			"message": "please, set up level log",
		})
		return
	}

	q := bson.M{}
	if service := r.URL.Query().Get("service"); service != "" {
		q["service"] = service
	}
	if from := r.URL.Query().Get("from"); from != "" {
		f, err := time.ParseInLocation("20060102_150405", from, shared.Location())
		if err == nil {
			q["timestamp"] = bson.M{
				"$gte": f,
			}
		}

	}
	col, ses := mongo.MongoCol(level)
	defer ses.Close()
	events := make([]types.LogMessage, 0)
	res := col.Find(q)
	if _, err := res.Count(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   true,
			"message": "error: " + err.Error(),
		})
		return
	}
	res.All(&events)

	if format := r.URL.Query().Get("format"); format != "" {
		switch format {
		case "csv":
			retVal := [][]string{}
			for _, e := range events {
				retVal = append(retVal, []string{
					e.Timestamp.Format("2006-01-02 15:04:05"), e.Service, "[" + strings.Join(e.Tags, ",") + "]", e.Message,
				})
			}
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			csv.NewWriter(w).WriteAll(retVal)
			return
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}
