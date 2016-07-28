<html>
    <head>
        <script>
            function redirect(){
                window.location.href = "http://{{.ReturnTo}}";         
            }
        </script>
    </head>
    <body onload="redirect()">
        <p>Please wait...</p>
        {{ $token := .Token }}
        {{ range $domain := .Domains }}
            <img src="{{$domain}}?token={{$token}}" style="display:none;"/>
        {{end}}
    </body>
</html>
