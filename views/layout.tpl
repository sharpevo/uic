<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>Login Portal</title>

        <link href="/static/css/bootstrap.min.css" rel="stylesheet">
        <link href="/static/css/bootstrap-watch.min.css" rel="stylesheet">
        <link href="/static/css/main.css" rel="stylesheet">
        <link href="/static/css/font-awesome.min.css" rel="stylesheet">

        <script src="/static/js/jquery-1.11.2.min.js"></script>
        <script src="/static/js/bootstrap.min.js"></script>
        <script src="/static/js/main.js"></script>
    </head>

    <body>
        <nav class="navbar navbar-default nav-fixed-top" role="navigation">
            <div id="navbar" class="navbar-collapse collapse">
                <ul class="nav navbar-nav navbar-right">
                    <li><a href="{{urlfor "HomeController.Get"}}">Home</a></li>

              {{ if .UserInfo.Id }}
              <li class="dropdown">
                  <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">{{.UserInfo.Email}}<span class="caret"></span></a>
                  <ul class="dropdown-menu" role="menu">
                        {{if hasRole .UserInfo "admin"}}
                        <li><a href="http://accounts.igenetech.com/roles">roles</a></li>
                        {{end}}
                      <li><a href="http://accounts.igenetech.com/profile">profile</a></li>
                      <li><a href="http://accounts.igenetech.com/logout">Logout</a></li>
                  </ul>
                </li>

              {{else}}


              <li>
               <a href="{{urlfor "RegisterController.Get"}}">Register</a>
              </li>
              <li>
               <a href="{{urlfor "LoginController.Get"}}">Login</a>
              </li>
              {{ end }}

                </ul>
            </div>
        </nav>

        <div class="container">
            {{ if .flash.notice }}
                <div class="alert alert-info alert-dismissible" role="alert">
                    <button type="button" class="close" data-dismiss="alert"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    {{.flash.notice}}
                </div>
            {{end}}

            {{ if .flash.error }}
                <div class="alert alert-danger alert-dismissible" role="alert">
                    <button type="button" class="close" data-dismiss="alert"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    {{.flash.error}}
                </div>
            {{end}}

            {{ if .flash.alert }}
                <div class="alert alert-warning alert-dismissible" role="alert">
                    <button type="button" class="close" data-dismiss="alert"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    {{.flash.warning}}
                </div>
            {{end}}
            {{.LayoutContent}}
        </div>
        <hr>

        <footer class="footer-text text-center">
            <ul class="list-unstyled">
                <li>
                    <ul class="list-inline">
                        <li>
                            &copy;2015
                        </li>
                        <li>
                            艾吉泰康生物科技（北京）有限公司 版权所有
                        </li>
                        <li>
                            京ICP备15025054号-2
                        </li>
                    </ul>
                </li>
                <li class="phone">
                    <ul class="list-inline">
                        <li>
                            <span class="glyphicon glyphicon-earphone"></span> : 010-84097967
                        </li>
                        <li>
                            <a href="mailto:market@igenetech.com" title="Contact us!">
                                <span class="glyphicon glyphicon-envelope"></span>
                                : market@<span style="display:none;">null</span>igenetech.com
                            </a>
                        </li>
                    </ul>
                </li>
            </ul>
        </footer>
    </body>
</html>
