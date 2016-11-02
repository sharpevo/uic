<a href="#" data-toggle="modal" data-id="" data-target="#createApp">
	Create
</a>

<table class="table table-hover table-striped">
    <thead>
        <tr>
            <th>Name</th>
            <th>Domain</th>
            <th>Remark</th>
            <th>Enabled</th>
        </tr>
    </thead>
    <tbody>
        {{range $index, $app := .Apps}}
        <tr data-id="{{$app.Id.Hex}}">
            <td>
                <a href="#" data-toggle="modal"
                    data-id="{{$app.Id.Hex}}"
                    data-name="{{$app.Name}}"
                    data-domain="{{$app.Domain}}"
                    data-remark="{{$app.Remark}}"
                    data-enabled="{{$app.Enabled}}"
                    data-target="#editApp">
                    {{$app.Name}}
                </a>
            </td>
            <td>
                {{$app.Domain}}
            </td>
            <td>
                {{$app.Remark}}
            </td>
            <td>
				{{$app.Enabled}}
                <a href="#" data-toggle="modal" data-id="{{$app.Id.Hex}}" data-name="{{$app.Name}}" data-target="#deleteApp" class="pull-right">
                    delete
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>
</table>

<div class="modal fade" id="createApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Add App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="appName" id="appName" class="form-control" value="" placeholder="targetseq" tabindex="1">
                    </div>
                    <div class="form-group">
                        <input type="text" name="appDomain" id="appDomain" class="form-control" value="" placeholder="www.targetseq.com" tabindex="2">
                    </div>
                    <div class="form-group">
                        <input type="text" name="appRemark" id="appRemark" class="form-control" value="" placeholder="notes..." tabindex="3">
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Add" class="btn btn-info" tabindex="4">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

<div class="modal fade" id="deleteApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Delete App</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <div class="modal-body">
                    <input type="hidden" name="appId" id="appId" class="form-control" value="">
                    <p>Delete app '<span id="appName"></span>' ?</p>
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
        $appId = $(event.target).data('id');
        $appName = $(event.target).data('name');
        $(this).find("#appId").val($appId);
        $(this).find("#appName").html($appName);
    });
});
</script>

<div class="modal fade" id="editApp" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="false">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel">Edit</h4>
            </div>
            <form action="{{urlfor "AppController.Post"}}" method="POST">
                <input type="hidden" name="appId" id="appId" class="form-control" value="">
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" name="appName" id="appName" class="form-control" value="" placeholder="targetseq" tabindex="1">
                    </div>
                    <div class="form-group">
                        <input type="text" name="appDomain" id="appDomain" class="form-control" value="" placeholder="www.targetseq.com" tabindex="2">
                    </div>
                    <div class="form-group">
                        <input type="text" name="appRemark" id="appRemark" class="form-control" value="" placeholder="notes..." tabindex="3">
                    </div>
                    <div class="form-group">
                        <input type="checkbox" name="appEnabled" id="appEnabled" tabindex="4">Enabled
                    </div>
                </div>
                <div class="modal-footer">
                    <input type="submit" value="Update" class="btn btn-info" tabindex="5">
                </div>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
<script>
$(function(){
    $('#editApp').on('show.bs.modal', function() {
        $appId = $(event.target).data('id');
        $appName = $(event.target).data('name');
        $appDomain = $(event.target).data('domain');
        $appRemark = $(event.target).data('remark');
        $appEnabled = $(event.target).data('enabled');
        $(this).find("#appId").val($appId);
        $(this).find("#appName").val($appName);
        $(this).find("#appDomain").val($appDomain);
        $(this).find("#appRemark").val($appRemark);
        $(this).find("#appEnabled").prop("checked", $appEnabled);
    });
});
</script>
