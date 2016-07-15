
<div id="login-overlay" class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
            <h4 class="modal-title" id="myModalLabel">Sign up to iGeneTech.com</h4>
             <!--<img class="pull-right" style="height: 30px; margin-top: -30px;" src="/static/img/logo.png" alt="iGeneTech Logo" />-->
        </div>
        <div class="modal-body">

<!--<div class="col-xs-12 col-sm-8 col-md-8 col-sm-offset-2 col-md-offset-2">-->
    <form role="form" action="{{urlfor "RegisterController.Post"}}" method="POST">
        <!--<h2>Please Sign Up <small>It's free and always will be.</small></h2>
        <hr/>-->
        <div class="form-group">
            <input type="text" name="name" id="name" class="form-control" value="{{.Name}}" placeholder="User Name" tabindex="3">
        </div>
        <div class="form-group">
            <input type="email" name="email" id="email" class="form-control" value="{{.Email}}" placeholder="Email Address" tabindex="4">
        </div>
        <div class="row">
            <div class="col-xs-12 col-sm-6 col-md-6">
                <div class="form-group">
                    <input type="password" name="password" id="password" class="form-control" placeholder="Password" tabindex="5">
                </div>
            </div>
            <div class="col-xs-12 col-sm-6 col-md-6">
                <div class="form-group">
                    <input type="password" name="password_confirmation" id="password_confirmation" class="form-control" placeholder="Confirm Password" tabindex="6">
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col-xs-12 col-sm-6 col-md-6">
                <div class="form-group">
                    {{create_captcha}}
                </div>
            </div>
            <div class="col-xs-12 col-sm-6 col-md-6">
                <div class="form-group" style="margin-top:1em;">
                    <div><small>If you can't read the captcha,</small></div>
                    <div><small>try another one by clicking on it.</small></div>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col-xs-12 col-sm-12 col-md-12">
                <div class="form-group">
                    <input name="captcha" type="text" class="form-control" placeholder="Enter the characters you see" tabindex="7" style="margin-top:1em;">
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col-xs-12 col-sm-12 col-md-12">
                <div class="form-group">
                    <small>Clicking "Register" button means that you agree to the <a href="#" data-toggle="modal" data-target="#agreement">iGeneTech Services Agreement</a>.</small>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col-xs-12 col-md-6"><a href="{{urlfor "LoginController.Get"}}" class="btn btn-success btn-block">Login</a></div>
            <div class="col-xs-12 col-md-6"><input type="submit" value="Register" class="btn btn-info btn-block" tabindex="8"></div>
        </div>
    </form>
<!--</div>-->

</div><!-- modal body -->
</div><!-- modal content -->
</div><!-- modal dialog -->

<!-- Modal -->
<div class="modal fade" id="agreement" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">iGeneTech Services Agreement</h4>
            </div>
            <div class="modal-body">
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
                <p>I WILL NOT BUY ANY KITS OF AGILENT</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-primary" data-dismiss="modal">I Agree</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

