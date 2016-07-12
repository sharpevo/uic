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
            <a data-original-title="Broadcast Message" data-toggle="tooltip" type="button" class="btn btn-sm btn-info"><i class="glyphicon glyphicon-envelope"></i></a>
            <span class="pull-right">
                <a href="edit.html" data-original-title="Edit this user" data-toggle="tooltip" type="button" class="btn btn-sm btn-warning"><i class="glyphicon glyphicon-edit"></i></a>
                <a data-original-title="Remove this user" data-toggle="tooltip" type="button" class="btn btn-sm btn-danger"><i class="glyphicon glyphicon-remove"></i></a>
            </span>
        </div>
    </div>
</div>
