<div id="animate">
    <div id="animate-item">
        <h1 id="targetseq" onclick="location.href='http://targetseq.igenetech.com'">
            TargetSeq
        </h1>
        <h4 id="targetseqFeature">
            A fast, flexible oligonucleotides design tool for small non-coding RNAs.
        </h4>
    </div>
    <div id="animate-item">
        <h1 id="multipseq" onclick="location.href='http://multipseq.igenetech.com'">
            MultipSeq
        </h1>
        <h4 id="multipseqFeature">
            A fast and versatile program for designing CRISPR sgRNA and evaluating potential off-target cleavage sites.
        </h4>
    </div>
    <div id="animate-item">
        <h1 id="primerqc" onclick="location.href='http://prmerqc.igenetech.com'">
            PrimerQC
        </h1>
        <h4 id="primerqcFeature">
            Analysis for PCR primer quality.
        </h4>
    </div>

    <div id="animate-item">
        <h1 id="crispr" onclick="location.href='http://crispr.igenetech.com'">
            CRISPR Design
        </h1>
        <h4 id="crisprFeature">
            A fast and versatile program for designing CRISPR sgRNA and evaluating potential off-target cleavage sites.
        </h4>
    </div>

    <div id="animate-item">
        <h1 id="srna" onclick="location.href='http://srnaprimer.igenetech.com'">
            sRNAPrimer
        </h1>
        <h4 id="srnaFeature">
            A fast, flexible oligonucleotides design tool for small non-coding RNAs.
        </h4>
    </div>
</div>
<style>
h1{
    color: #11aae4;
    visibility:hidden;
    padding-left: 100px;
    cursor:pointer;
    cursor:hand;
}

h4{
    color: #111111;
    visibility:hidden;
    padding-left: 163px;
}

#animate-item{
    padding-top: 25px;
    padding-bottom: 25px;
}
</style>
<script>
$('h1').textillate({
    sync:true,
    type:'char',
    autoStart:false,
    in:{
        effect:'fadeInLeft',
    },
    out:{
        effect:'fadeOutLeft',
    }
})
$('h4').textillate({
    sequence:true,
    type:'word',
    autoStart:false,
    in:{
        effect:'fadeInRight',
    },
    out:{
        effect:'fadeOutRight',
    }
})


$('#targetseq').textillate('in')
$('#targetseq').on('inAnimationEnd.tlt', function () {
    $('#targetseqFeature').textillate('in')
});
$('#targetseqFeature').on('inAnimationEnd.tlt', function () {
    $('#multipseq').textillate('in')
})


$('#multipseq').on('inAnimationEnd.tlt', function () {
    $('#multipseqFeature').textillate('in')
})
$('#multipseqFeature').on('inAnimationEnd.tlt', function () {
    $('#primerqc').textillate('in')
})


$('#primerqc').on('inAnimationEnd.tlt', function () {
    $('#primerqcFeature').textillate('in')
})
$('#primerqcFeature').on('inAnimationEnd.tlt', function () {
    $('#crispr').textillate('in')
})


$('#crispr').on('inAnimationEnd.tlt', function () {
    $('#crisprFeature').textillate('in')
})
$('#crisprFeature').on('inAnimationEnd.tlt', function () {
    $('#srna').textillate('in')
})


$('#srna').on('inAnimationEnd.tlt', function () {
    $('#srnaFeature').textillate('in')
})
$('#srnaFeature').on('inAnimationEnd.tlt', function () {
    $('#targetseq').textillate('out')
    $('#targetseqFeature').textillate('out')

    $('#multipseq').textillate('out')
    $('#multipseqFeature').textillate('out')

    $('#primerqc').textillate('out')
    $('#primerqcFeature').textillate('out')

    $('#crispr').textillate('out')
    $('#crisprFeature').textillate('out')

    $('#srna').textillate('out')
    $('#srnaFeature').textillate('out')
})

$('#srnaFeature').on('outAnimationEnd.tlt', function () {
    $('#targetseq').textillate('in')
})


</script>
