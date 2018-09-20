package main
import (
  "fmt"
   "log"
   "net/http"
  "html/template"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  //"os"
   
  
)

type Cricket struct {
      Name string
       Gender string
	Age string
	Location string
	Applied_as string
	Style string
	Applied_on string
        Trial_on_date string
        Trial_on_time string
        
        
}

func process(w http.ResponseWriter, r *http.Request) { 

   //fmt.Println("hii") 
  t, _ := template.ParseFiles("all-players.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("all")
 
  err = c.Find(nil).Sort("-$natural").All(&results)
  
  t.Execute(w, results)

 //err = c.Find(bson.M{"role": "Bowler"}).All(&bowl)
 //t.Execute(w, bowl)
  
} 



func shortdb(w http.ResponseWriter, r *http.Request) { 
      if r.Method !="POST" {
      http.Redirect(w,r,"/",http.StatusSeeOther) 
 	return 
 	}
 t, _ := template.ParseFiles("all-players.html")
 

var results []Cricket
err :=r.ParseForm()
 
//var value=r.PostFormValue("value_1")

contact:=make(map[string]string)
contact["name"]=r.FormValue("name")
contact["gender"]=r.FormValue("gender")
contact["age"]=r.PostFormValue("age") 
contact["location"]=r.PostFormValue("location")
contact["applied_as"]=r.PostFormValue("applied_as")
contact["style"]=r.PostFormValue("style")
contact["applied_on"]=r.PostFormValue("applied_on")

  //fmt.Println(contact)
  //fmt.Println("SHorted")
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
 
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("all")
  d :=session.DB("dumb").C("shortlist")
  
  // err = c.Find(nil).Sort("-$natural").All(&results)
  //t.Execute(w,results)
 
  //err = c.Find(bson.M{"role": "Bowler"}).All(&bowl)
var count int
count, err = d.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()

//fmt.Print(count)
err = d.Find(nil).All(&results)
if count>0 {
   //process(w,r)
    log.Printf("Already Exist")
   t.Execute(w,nil)
} else{
err = d.Insert(&contact)
log.Printf("Inserted")
t.Execute(w,results)
}

 
err=c.Remove(bson.M{"name": contact["name"]})
if err != nil {
    fmt.Printf("remove fail %v\n", err)
    //os.Exit(1)
   }





} 



func rejectdb(w http.ResponseWriter, r *http.Request) { 
       if r.Method !="POST" {
      http.Redirect(w,r,"/",http.StatusSeeOther) 
 	return 
 	}
 t, _ := template.ParseFiles("all-players.html")
 

var results []Cricket
err :=r.ParseForm()
 
//var value=r.PostFormValue("value_1")

contact:=make(map[string]string)
contact["name"]=r.FormValue("name")
contact["gender"]=r.FormValue("gender")
contact["age"]=r.PostFormValue("age") 
contact["location"]=r.PostFormValue("location")
contact["applied_as"]=r.PostFormValue("applied_as")
contact["style"]=r.PostFormValue("style")
contact["applied_on"]=r.PostFormValue("applied_on")

 // fmt.Println(contact)
  
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
 
  session.SetMode(mgo.Monotonic, true)
  c :=session.DB("dumb").C("all")
  d :=session.DB("dumb").C("reject")
  
  f :=session.DB("dumb").C("interview")
  g :=session.DB("dumb").C("shortlist")
  // err = c.Find(nil).Sort("-$natural").All(&results)
  //t.Execute(w,results)
 
  
var count int
count, err = d.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()

//fmt.Print(count)

if count>0 {
   
    //log.Printf("Already Exist")
   t.Execute(w,nil)
} else{
err = d.Insert(&contact)
//log.Printf("Rejected")
t.Execute(w,results)
}



err=c.Remove(bson.M{"name": contact["name"]})
if err != nil {
    //fmt.Printf("remove fail %v\n", err)
    
   }




err=f.Remove(bson.M{"name": contact["name"]})
if err != nil {
    //fmt.Printf("remove fail %v\n", err)
    
   }


err=g.Remove(bson.M{"name": contact["name"]})
if err != nil {
    //fmt.Printf("remove fail %v\n", err)
    
   }

} 





func shortlst(w http.ResponseWriter, r *http.Request) { 

   //fmt.Println("hii") 
  t, _ := template.ParseFiles("shortplay.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("shortlist")
   
    
  err = c.Find(nil).Sort("-$natural").All(&results)
  //fmt.Println(results)
  t.Execute(w, results)

  
} 






func rejected(w http.ResponseWriter, r *http.Request) { 

   //fmt.Println("hii") 
  t, _ := template.ParseFiles("rejectplay.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("reject")
   
    
  err = c.Find(nil).Sort("-$natural").All(&results)
  //fmt.Println(results)
  t.Execute(w, results)

  
} 







func index(w http.ResponseWriter, r *http.Request) { 


 t, _ := template.ParseFiles("index.html")


  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("all")
  d := session.DB("dumb").C("interview")
  e := session.DB("dumb").C("shortlist")
  f :=session.DB("dumb").C("select")
  err = c.Find(nil).Sort("-$natural").Limit(5).All(&results)
  //fmt.Println(results)
  t.Execute(w, results)
  
  //err = d.Find(nil).Sort("-$natural").Limit(5).All(&latest)
  // u.Execute(w, latest)
  




 var total_applied int
 var total_interview int
 var total_shortlist int
 var total_selected int

total_applied, err = c.Find(nil).Count()
total_interview, err = d.Find(nil).Count()
total_shortlist,err=e.Find(nil).Count()
total_selected,err=f.Find(nil).Count()
  varmap := map[string]interface{}{
            "total": total_applied,
            "interview": total_interview,
	    "short": total_shortlist,
	    "select": total_selected,	
        }
   t.ExecuteTemplate(w, "index", varmap)


  var interview []Cricket

   err=d.Find(nil).All(&interview)    
t.ExecuteTemplate(w,"inter",interview)
  
}







func verifydb(w http.ResponseWriter, r *http.Request) { 
       if r.Method !="POST" {
      http.Redirect(w,r,"/",http.StatusSeeOther) 
 	return 
 	}
 t, _ := template.ParseFiles("all-players.html")
 

var results []Cricket
err :=r.ParseForm()
 
//var value=r.PostFormValue("value_1")

contact:=make(map[string]string)
contact["name"]=r.FormValue("name")
contact["gender"]=r.FormValue("gender")
contact["age"]=r.PostFormValue("age") 
contact["location"]=r.PostFormValue("location")
contact["applied_as"]=r.PostFormValue("applied_as")
contact["style"]=r.PostFormValue("style")
contact["applied_on"]=r.PostFormValue("applied_on")

 // fmt.Println(contact)
  
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
 
  session.SetMode(mgo.Monotonic, true)
  c :=session.DB("dumb").C("all")
  d :=session.DB("dumb").C("verified")
  // err = c.Find(nil).Sort("-$natural").All(&results)
  //t.Execute(w,results)
 
  //err = c.Find(bson.M{"role": "Bowler"}).All(&bowl)
var count int
count, err = d.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()

//fmt.Print(count)
//err = d.Find(nil).All(&results)
if count>0 {
   //process(w,r)
    log.Printf("Already Exist")
   t.Execute(w,nil)
} else{
err = d.Insert(&contact)
//log.Printf("Rejected")
t.Execute(w,results)
}



err=c.Remove(bson.M{"name": contact["name"]})
if err != nil {
    fmt.Printf("remove fail %v\n", err)
    
   }


} 

func verified(w http.ResponseWriter, r *http.Request) { 


 t, _ := template.ParseFiles("varified.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("verified")
   
    
  err = c.Find(nil).Sort("-$natural").All(&results)
  //fmt.Println(results)
  t.Execute(w, results)


}


func trialdb(w http.ResponseWriter, r *http.Request) { 
       if r.Method !="POST" {
      http.Redirect(w,r,"/",http.StatusSeeOther) 
 	return 
 	}

 


err :=r.ParseForm()
 


contact:=make(map[string]string)
contact["name"]=r.FormValue("name")
contact["gender"]=r.FormValue("gender")
contact["age"]=r.PostFormValue("age") 
contact["location"]=r.PostFormValue("location")
contact["applied_as"]=r.PostFormValue("applied_as")
contact["style"]=r.PostFormValue("style")
contact["applied_on"]=r.PostFormValue("applied_on")
contact["trial_on_date"]=r.PostFormValue("trial_on_date")
contact["trial_on_time"]=r.PostFormValue("trial_on_time")

 
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
 
  session.SetMode(mgo.Monotonic, true)

  e :=session.DB("dumb").C("interview")
  f :=session.DB("dumb").C("shortlist")
  
var count int
count, err = e.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()


if count>0 {
   
    fmt.Printf("Already Exist")
  
} else{
err = e.Insert(&contact)
}



err=f.Remove(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]})
if err != nil {
    //fmt.Printf("remove fail %v\n", err)
    
   }


} 



func trialed(w http.ResponseWriter, r *http.Request) { 

    
  t, _ := template.ParseFiles("trialplay.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("interview")
   
    
  err = c.Find(nil).Sort("-$natural").All(&results)
  
  t.Execute(w, results)

  
} 


func selectdb(w http.ResponseWriter, r *http.Request) { 
       if r.Method !="POST" {
      http.Redirect(w,r,"/",http.StatusSeeOther) 
 	return 
 	}

 


err :=r.ParseForm()
 


contact:=make(map[string]string)
contact["name"]=r.FormValue("name")
contact["gender"]=r.FormValue("gender")
contact["age"]=r.PostFormValue("age") 
contact["location"]=r.PostFormValue("location")
contact["applied_as"]=r.PostFormValue("applied_as")
contact["style"]=r.PostFormValue("style")
contact["applied_on"]=r.PostFormValue("applied_on")


 
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
 
  session.SetMode(mgo.Monotonic, true)

  e :=session.DB("dumb").C("interview")
  f :=session.DB("dumb").C("select")
  g :=session.DB("dumb").C("reject")
var count int
count, err = f.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()
var rejcount int
rejcount,err=g.Find(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]}).Count()

if count>0 {
   
    fmt.Printf("Already Exist")
  
} else{
err = f.Insert(&contact)
}

if rejcount>0  {
  err=g.Remove(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]})
} else{
fmt.Println("Congratulations")
}

err=e.Remove(bson.M{"name": contact["name"],"gender":contact["gender"],"age":contact["age"]})
if err != nil {
   // fmt.Printf("remove fail %v\n", err)
    
   }




} 


func selected(w http.ResponseWriter, r *http.Request) { 

    
  t, _ := template.ParseFiles("selectplay.html")

  
  var results []Cricket
 
   
  session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c := session.DB("dumb").C("select")
   
    
  err = c.Find(nil).Sort("-$natural").All(&results)
 // fmt.Println(results)
  t.Execute(w, results)

  
}


func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  fmt.Println("Started..")
  
  http.HandleFunc("/process", process)
  http.HandleFunc("/short", shortdb)
  http.HandleFunc("/reject",rejectdb)
  http.HandleFunc("/shortplay", shortlst)
  http.HandleFunc("/rejectplay", rejected)
  http.HandleFunc("/index", index)
  http.HandleFunc("/verify", verifydb)
  http.HandleFunc("/verifyplay", verified)
  http.HandleFunc("/trial", trialdb)
  http.HandleFunc("/trialplay", trialed)
  http.HandleFunc("/select", selectdb) 
  http.HandleFunc("/selectplay", selected) 
  
 
  server.ListenAndServe()
}
