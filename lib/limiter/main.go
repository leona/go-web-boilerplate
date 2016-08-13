package limiter

import (
    "github.com/jaredfolkins/badactor"
    "time"
    "github.com/neoh/go-web-boilerplate/lib/logger"
)

var St *badactor.Studio

func init() {
    St = badactor.NewStudio(256)
  
    // create and add rule
    ru := &badactor.Rule{
      Name:        "NOT_GET",
      Message:     "Too many post requests",
      StrikeLimit: 10,
      ExpireBase:  time.Second * 15, 
      Sentence:    time.Minute * 1, 
    }
    
    // add the rule to the stack
    St.AddRule(ru)

    
    // creates the Directors who act as the Buckets in our sharding cache
    err := St.CreateDirectors(256)
    if err != nil {
      logger.Fatal("Create director fail")
    }
    
    //poll duration 
    dur := time.Minute * time.Duration(60)
    // Start the reaper
    St.StartReaper(dur)
}