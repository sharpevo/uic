<div id="login-overlay" class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
            <h4 class="modal-title" id="myModalLabel">Recover your account <small> Enter your email to get the reset mail</small></h4>
        </div>
        <div class="modal-body">
            <form role="form" class="form" action="{{urlfor "ForgotPasswdController.Post"}}" method="post">
                <div class="row">
                    {{ .xsrfdata }}
                    <div class="col-sm-12 col-md-12">
                        <div class="form-group">
                            <input type="email" id="email" name="email" value="{{.Email}}" class="form-control text-left" placeholder="Enter your email address" tabindex="1" />
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-xs-6 col-sm-6 col-md-6">
                        <div class="form-group">
                            <input name="captcha" type="text" class="form-control" placeholder="Enter the characters you see" tabindex="2" style="margin-top:1em;">
                        </div>
                    </div>
                    <div class="col-xs-6 col-sm-6 col-md-6">
                        <div class="form-group">
                            {{create_captcha}}
                        </div>
                    </div>
                </div>

                <div class="row">
                    <div class="col-xs-12 col-md-6">
                        <a href="{{urlfor "LoginController.Get"}}" class="btn btn-success btn-block">Login</a>
                    </div>
                    <div class="col-xs-12 col-md-6">
                        <input class="btn btn-info btn-block" name="commit" type="submit" value="Send Reset Mail"/>
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
