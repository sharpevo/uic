<table class="table table-hover table-striped">
    <thead>
        <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Roles</th>
            <th>UserOperation</th>
        </tr>
    </thead>
    <tbody>
        {{range $index, $user := .UserList}}
        <tr data-id="{{$user.Id.Hex}}">
            <td>
                {{$user.Name}}<br/>
                {{range $k, $v := $user.Apps}}
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-appname="{{SpaceToDot $k}}" data-target="#deleteApp" class="badge" style="margin-top:5px;"> {{SpaceToDot $k}}</a><br/>
                {{end}}
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-target="#createApp">
                    <span class="glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
                </a>
            </td>
            <td>{{$user.Email}}</td>
            <td>
                {{range $k, $v := $user.Roles}}
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-rolename="{{$k}}" data-target="#deleteRole"> {{$k}} </a><br/>
                {{end}} 
                <a href="#" data-toggle="modal" data-id="{{$user.Id.Hex}}" data-target="#createRole">
                    <span class="glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
                </a>
            </td>
            <td>Edit</td>
        </tr>
        {{end}}
    </tbody>
</table>

<!-- Modal -->
<div class="modal fade" id="createRole" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Add Role</h4>
            </div>
            <form action="{{urlfor "RoleController.Post"}}" method="POST">
                <input type="hidden" name="userId" id="userId" class="form-control" value="" placeholder="Enter the role name">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="roleName" id="roleName" class="form-control" value="" placeholder="Enter the role name" tabindex="1">
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Add" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $("#createRole").on('show.bs.modal', function() {
        $userId = $(event.target).closest('tr').data('id');
        $(this).find("#userId").val($userId);
        console.log($userId)
        setTimeout(function() {
            $('#roleName').focus();
        }, 500);
    });
});
</script>

<div class="modal fade" id="deleteRole" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Delete Role</h4>
            </div>
            <form action="{{urlfor "RoleController.Post"}}" method="POST">
                <div class="modal-body">
                    <input type="hidden" name="userId" id="userId" class="form-control" value="" placeholder="Enter the role name">
                    <input type="hidden" name="roleName" id="roleName" class="form-control" value="" placeholder="Enter the role name">
                    <p>Delete role '<span id="roleNameShown"></span>' ?</p>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Confirm" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $('#deleteRole').on('show.bs.modal', function() {
        $userId = $(event.target).data('id');
        $roleName = $(event.target).data('rolename');
        $(this).find("#userId").val($userId);
        $(this).find("#roleName").val($roleName);
        $(this).find("#roleNameShown").html($roleName);
    });
});
</script>

<div class="modal fade" id="createApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Add App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <input type="hidden" name="userId" id="userId" class="form-control" value="">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="appName" id="appName" class="form-control" value="" placeholder="Enter the app name" tabindex="1">
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Add" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $("#createApp").on('show.bs.modal', function() {
        $userId = $(event.target).closest('tr').data('id');
        $(this).find("#userId").val($userId);
        console.log($userId)
        setTimeout(function() {
            $('#appName').focus();
        }, 500);
    });
});
</script>

<div class="modal fade" id="deleteApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Delete App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <div class="modal-body">
                    <input type="hidden" name="userId" id="userId" class="form-control" value="">
                    <input type="hidden" name="appName" id="appName" class="form-control" value="">
                    <p>Delete app '<span id="appNameShown"></span>' ?</p>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Confirm" class="btn btn-info" tabindex="2">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $('#deleteApp').on('show.bs.modal', function() {
        $userId = $(event.target).data('id');
        $appName = $(event.target).data('appname');
        $(this).find("#userId").val($userId);
        $(this).find("#appName").val($appName);
        $(this).find("#appNameShown").html($appName);
    });
});
</script>
