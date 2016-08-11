package settings

type keys struct {
    Csrf string
}

type errors struct {
    IntervalServer []byte
    NotFound []byte
}

var Keys *keys
var Errors *errors

func init() {
    Keys = &keys{
        Csrf: "csrf_token",
    }
    
    Errors = &errors{
        IntervalServer: []byte("CUSTOM 500 INTERNAL SERVER ERROR PAGE"),
        NotFound: []byte("CUSTOM 404 NOT FOUND ERROR PAGE"),
    }
}