package worker

var MaxLength int64 = 10000

type ActionCollection struct {
	actions []Action
}

//func payloadHandler(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != "POST" {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		return
//	}
//
//	// Read the body into a string for json decoding
//	var content = &ActionCollection{}
//	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// Go through each payload and queue items individually to be posted to S3
//	for _, ac := range content.actions {
//
//		// let's create a job with the payload
//		work := Job{AC: ac}
//
//		// Push the work onto the queue.
//		JobQueue <- work
//	}
//
//	w.WriteHeader(http.StatusOK)
//}

func Add() {
	j := Job{AC: Action{Name: "zj", Age: 11}}
	JobQueue <- j
}
