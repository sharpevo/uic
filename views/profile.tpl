<div id="profile-overlay" class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
            <h4 class="modal-title" id="myModalLabel">Profile</h4>
        </div>
        <div class="modal-body">
            <div class="row">
                <div class="col-md-3 col-lg-3 " align="center">
                    <!--<img alt="User Pic" src="http://babyinfoforyou.com/wp-content/uploads/2014/10/avatar-300x300.png" class="img-circle img-responsive">-->
                    <img alt="User Pic" src="https://tse3.mm.bing.net/th?id=OIP.Me6d1b1332b4e350d80c9932d0bfdfe6ao0&pid=15.1" class="img-circle img-responsive"> 
                </div>
                <div class=" col-md-9 col-lg-9 "> 
                    <table class="table">
                        <tbody>
                            <tr>
                                <td>UserName:</td>
                                <td>{{.UserInfo.Name}}</td>
                            </tr>
                            <tr>
                                <td>Email:</td>
                                <td>{{.UserInfo.Email}}</td>
                            </tr>
                            <tr>
                                <td>Password:</td>
                                <td>******</td>
                            </tr>
                            <tr>
                                <td>Register Date:</td>
                                <td>{{Format .User.DateCreated}}</td>
                            </tr>
                            <tr>
                                <td>Last login:</td>
                                <td>{{Format .User.DateLastLogin}}</td>
                            </tr>
                        </tbody>
                    </table>
                    <span class="label label-pill label-primary">PrimerQC</span>
                    <span class="label label-pill label-primary">sRNAPrimer</span>
                    <span class="label label-primary">CRISPR</span>
                    <span class="label label-default">MultipSeq</span>
                    <span class="label label-default">TargetSeq</span>
                </div>
            </div>
        </div>
        <div class="panel-footer">
            <a href="#" data-toggle="modal" data-target="#deleteUser" class="btn btn-sm btn-danger" style="visibility:hidden"><i class="glyphicon glyphicon-remove"></i></a>
            <span class="pull-right">
                 <a href="#" data-toggle="modal" data-target="#editProfile" class="btn btn-sm btn-warning">Edit</a>
            </span>
        </div>
    </div>
</div>

<!-- Modal -->
<div class="modal fade" id="editProfile" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Edit Profile</h4>
            </div>
            <form role="form" action="{{urlfor "ProfileController.Post"}}" method="POST">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="name" id="name" class="form-control" value="{{.UserInfo.Name}}" placeholder="User Name" tabindex="3">
                    </div>
                    <div class="form-group">
                        <input type="email" name="email" id="email" class="form-control" value="{{.UserInfo.Email}}" placeholder="Email Address" tabindex="4">
                    </div>

                    <div class="form-group">
                        <small>Leave the passwords blank if you don't want to change it.</small>
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
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Submit" class="btn btn-primary"/>
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
