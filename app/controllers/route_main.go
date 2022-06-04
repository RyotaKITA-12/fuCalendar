package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        generateHTML(w, nil, "layout", "public_navbar", "top")
    } else {
        http.Redirect(w, r, "/todos", 302)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/", 302)
    } else {
        user, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        todos, _ := user.GetTodosByUser()
        user.Todos = todos
        generateHTML(w, user, "layout", "private_navbar", "index")
    }
}

func eventNew(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "event_new")
    }
}

func eventSave(w http.ResponseWriter, r *http.Request) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        err = r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        user, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        content := r.PostFormValue("content")
        location := r.PostFormValue("location")
        start_time := stringToTime(r.PostFormValue("start_time"))
        end_time := stringToTime(r.PostFormValue("end_time"))
        if err := user.CreateEvent(content, location, start_time, end_time); err != nil {
            log.Println(err)
        }
        http.Redirect(w, r, "/events", 302)
    }
}

func stringToTime(str string) time.Time {
    var layout = "2022-01-01 10:00:00"
    t, _ := time.Parse(layout, str)
    return t
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        _, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        t, err := models.GetTodo(id)
        if err != nil {
            log.Println(err)
        }
        generateHTML(w, t, "layout", "private_navbar", "todo_edit")
    }
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        user, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        content := r.PostFormValue("content")
        t := &models.Todo{ID: id, Content: content, UserID: user.ID}
        if err := t.UpdateTodo(); err != nil {
            log.Println(err)
        }
        http.Redirect(w, r, "/todos", 302)
    }
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        _, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        t, err := models.GetTodo(id)
        if err != nil {
            log.Println(err)
        }
        if err := t.DeleteTodo(); err != nil {
            log.Println(err)
        }
        http.Redirect(w, r, "/todos", 302)
    }
}
