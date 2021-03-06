<div id="login-overlay" class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
            <h4 class="modal-title" id="myModalLabel">Login to iGeneTech.com</h4>
        </div>
        <div class="modal-body">
            <div class="row">
                <div class="col-xs-6">
                    <div class="well">
                        <form id="loginForm" method="POST" action="{{urlfor "LoginController.Post"}}" novalidate="novalidate">
                            <div class="form-group">
                                <label for="email" class="control-label">Email</label>
                                <input type="text" class="form-control" id="emair" name="email" value="{{.Email}}" required="" title="Please enter you email" placeholder="example@gmail.com">
                                <input type="hidden" name="return_to" value="{{.ReturnTo}}">
                                <span class="help-block"></span>
                            </div>
                            <div class="form-group">
                                <label for="password" class="control-label">Password</label>
                                <input type="password" class="form-control" id="password" name="password" value="" required="" title="Please enter your password">
                                <span class="help-block"></span>
                            </div>
                            <div id="loginErrorMsg" class="alert alert-error hide">Wrong username og password</div>
                            <div class="checkbox">
                                <label>
                                    <input type="checkbox" name="remember" id="remember"> Remember me for a month.
                                </label>
                                <p class="help-block">(if this is a private computer)</p>
                            </div>
                            <button type="submit" class="btn btn-success btn-block">Login</button>
                            <a href="{{urlfor "ForgotPasswdController.Get"}}" class="btn btn-default btn-block">Help to login</a>
                        </form>
                    </div>
                </div>
                <div class="col-xs-6">
                    <p class="lead">Register now for <span class="text-success">FREE</span></p>
                    <ul class="list-unstyled" style="line-height: 2">
                        <li><span class="fa fa-check text-success"></span> CRISPR</li>
                        <li><span class="fa fa-check text-success"></span> PrimerQC</li>
                        <li><span class="fa fa-check text-success"></span> sRNAPrimer</li>
                        <li><span class="fa fa-check text-success"></span> TargetSeq<small> (only for registered user)</small></li>
                        <li><span class="fa fa-check text-success"></span> MultipSeq<small> (only for registered user)</small></li>
                        <li><a href="http://www.igenetech.com"><u>Read more</u></a></li>
                    </ul>
                    <p><a href="{{urlfor "RegisterController.Get"}}" class="btn btn-info btn-block {{if not .SignUpEnabled}}disabled{{end}}">Yes please, register now!</a></p>
                </div>
            </div>
        </div>
    </div>
</div>
