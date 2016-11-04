<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>User Information Center - iGeneTech</title>

        <link href="/static/css/bootstrap.min.css" rel="stylesheet">
        <link href="/static/css/bootstrap-watch.min.css" rel="stylesheet">
        <link href="/static/css/main.css" rel="stylesheet">
        <link href="/static/css/font-awesome.min.css" rel="stylesheet">

        <script src="/static/js/jquery-1.11.2.min.js"></script>
        <script src="/static/js/bootstrap.min.js"></script>
        <script src="/static/js/main.js"></script>

        <link href="/static/css/animate.min.css" rel="stylesheet">
        <script src="/static/js/jquery.lettering.js"></script>
        <script src="/static/js/jquery.textillate.js"></script>
        <style>
body{
    padding:70px 0 70px 0;
}
.customlink, .customlink:hover{
    text-decoration: none;
    color:#fff;
    margin-left: 10px;
}
        </style>
    </head>

    <body>
        <!--<img id="iGeneTechLogo" class="pull-right" src="/static/img/logo.png" alt="iGeneTech" height="40px"  style="margin-top: -8px"/></p>-->
        <div class="navbar navbar-default navbar-fixed-top" role="navigation">
            <div class="container-fluid">
                <a  class="navbar-brand"  rel="home"  href="http://{{.UICDomain}}"  title="iGeneTech">
                    User Information Center
                </a>
                <ul class="nav navbar-nav navbar-right">
                    {{ if .UserInfo.Id }}
                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">{{.UserInfo.Email}}<span class="caret"></span></a>
                        <ul class="dropdown-menu" role="menu">
                            <li><a href="http://{{.UICDomain}}/profile">Account Settings <span class="glyphicon glyphicon-cog pull-right"></span></a></li>
                            {{if hasRole .UserInfo "admin"}}
                            <li><a href="http://{{.UICDomain}}/roles">Manage Users <span class="glyphicon glyphicon-user pull-right"></span></a></li>
                            <li><a href="http://{{.UICDomain}}/apps">Manage Apps <span class="glyphicon glyphicon-th-large pull-right"></span></a></li>
                            <li class="divider"></li>
                            {{end}}
                            <li><a href="http://{{.UICDomain}}/logout">Logout <span class="glyphicon glyphicon-log-out pull-right"></span></a></li>
                        </ul>
                    </li>

                    {{else}}

                    {{if .SignUpEnabled}}
                    <li>
                        <a href="{{urlfor "RegisterController.Get"}}">Register</a>
                    </li>
                    {{end}}
                    <li>
                        <a href="{{urlfor "LoginController.Get"}}">Login</a>
                    </li>
                    {{ end }}
                </ul>
            </div>
        </div>

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

            {{ if .flash.warning }}
            <div class="alert alert-warning alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                {{.flash.warning}}
            </div>
            {{end}}
            {{.LayoutContent}}
        </div>

        <div class="navbar navbar-default navbar-fixed-bottom">
            <div class="container-fluid" style="margin-bottom: -30px;">
                <p class="navbar-text navbar-left">
                &copy;2016 艾吉泰康生物科技（北京）有限公司 版权所有 京ICP备15025054号-2
                </p>
                <p class="navbar-text navbar-right">
                    <a href="tel:4008190260" class="customlink">
                        Tel: 4008190260
                    </a>
                    <a href="mailto:market@igenetech.com" class="customlink">
                        Mail: market@<span style="display:none;">null</span>igenetech.com
                    </a>
                </p>
            </div>
        </div>
    </body>
</html>
