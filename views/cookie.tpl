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
        <img src="http://www.multipseq.com?token={{.Token}}" style="display:none;"/>
        <img src="http://www.targetseq.com?token={{.Token}}" style="display:none;"/>
        <img src="http://www.designhub.com?token={{.Token}}" style="display:none;"/>
    </body>
</html>
