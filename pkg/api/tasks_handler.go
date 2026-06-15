package api

import (
    "net/http"
    "todo/pkg/db"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
    search := r.URL.Query().Get("search")
    limit := 50
    var tasks []*db.Task
    var err error
    if search != "" {
        tasks, err = db.SearchTasks(search, limit)
    } else {
        tasks, err = db.Tasks(limit)
    }
    if err != nil {
        errorResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if tasks == nil {
        tasks = []*db.Task{}
    }
    writeJSON(w, map[string]interface{}{"tasks": tasks}, http.StatusOK)
}
