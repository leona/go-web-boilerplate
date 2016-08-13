package middleware

import (
    "fmt"
    "github.com/kataras/iris"
    "github.com/iris-contrib/middleware/secure"
    "github.com/neoh/go-web-boilerplate/lib/settings"
)

func init() {
    fmt.Println("Loading secure headers middleware")
    
    s := secure.New(secure.Options{
        AllowedHosts:            []string{ settings.Config.Host },                                                                                                                         
        //SSLRedirect:             true,     
        SSLTemporaryRedirect:    false,    
        SSLHost:                 settings.Config.Host,
        SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"},
        STSSeconds:              315360000,                                                                                                                                           
        STSIncludeSubdomains:    true,                                                                                                                                                
        STSPreload:              true,    
        ForceSTSHeader:          true,  
        FrameDeny:               true,    
        CustomFrameOptionsValue: "SAMEORIGIN",
        ContentTypeNosniff:      true,  
        BrowserXSSFilter:        true,
        ContentSecurityPolicy:   "default-src 'self'",   
        //PublicKey:               `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
        // PublicKey implements HPKP to prevent 
        //MITM attacks with forged certificates. Default is "".
        IsDevelopment: settings.Config.DevelopmentMode,
    })
    
    iris.UseFunc(func(c *iris.Context) {
        err := s.Process(c)

        if err != nil {
            return
        }
    
        c.Next()
    })
}
