
<div id="login-overlay" class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
            <h4 class="modal-title" id="myModalLabel">Reset Password</h4>
        </div>
        <div class="modal-body">
            <form role="form" class="form" action="{{urlfor "ResetPasswdController.Post"}}" method="post">
                <div class="row">
                    <div class="col-sm-12 col-md-12">
                        {{ .xsrfdata }}
                        <input type="hidden" id="token" name="token" value="{{.token}}" class="form-control text-left" />
                        <div class="form-group">
                            <input type="password" id="password" name="password" value="" class="form-control text-left" placeholder="Password (6-15)"/>
                        </div>
                        <div class="form-group">
                            <input type="password" id="password_confirmation" name="password_confirmation" value="" class="form-control text-left" placeholder="Confirm password" />
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-xs-12 col-md-6">
                        <a href="{{urlfor "LoginController.Get"}}" class="btn btn-success btn-block">Login</a>
                    </div>
                    <div class="col-xs-12 col-md-6">
                        <input class="btn btn-warning btn-block" name="commit" type="submit" value="Reset" />
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
