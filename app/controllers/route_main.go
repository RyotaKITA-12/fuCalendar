package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/RyotaKITA-12/fuCalendar.git/app/models"
)

func index(w http.ResponseWriter, r *http.Request) {
    sess, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/", 302)
    } else {
        user, err := sess.GetUserBySession()
        if err != nil {
            log.Println(err)
        }
        events, _ := user.GetEventsByUser()
        if err != nil {
            log.Println(err)
        }
        user.Events = events
        generateHTML(w, user, "layout", "private_navbar", "index")
    }
}

func search(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "search")
    }
}

func searchTime(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "search_time")
    }
}

func searchFriend(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "search_friend")
    }
}

func group(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "group")
    }
}

func groupSave(w http.ResponseWriter, r *http.Request) {
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
        // name := r.PostForm
        generateHTML(w, user, "layout", "private_navbar", "group")
    }
}

func schedule(w http.ResponseWriter, r *http.Request) {
    _, err := session(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", 302)
    } else {
        generateHTML(w, nil, "layout", "private_navbar", "schedule")
    }
}

func scheduleSave(w http.ResponseWriter, r *http.Request) {
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
        group_id, _ := strconv.Atoi(r.PostFormValue("group"))
        if err := user.CreateEvent(content, location, start_time, end_time, group_id); err != nil {
            log.Println(err)
        }
        http.Redirect(w, r, "/schedule", 302)
    }
}

func stringToTime(str string) time.Time {
    var layout = "2022-01-01 10:00:00"
    t, _ := time.Parse(layout, str)
    return t
}
