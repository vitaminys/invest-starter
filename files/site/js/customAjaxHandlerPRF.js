<<<<<<< HEAD
$( "#formVK" ).submit(function(event) {
  event.preventDefault();
  
  var $form = $( this ),
	vkidI = $form.find( "input[name='vkid']" ).val(),
	urlI = $form.attr( "action" );
	
  var posting = $.post( urlI, { vklink: vkidI })
  .always(function() {
	$("#loadVK").show();
	$(".vkDAE").attr("disabled");
  })
  .done(function() {
    $("#loadVK").hide();
	$(".vkDAE").removeAttr("disabled");
  })
  .fail(function() {
    $("#loadVK").hide();
	$(".vkDAE").removeAttr("disabled");
  });
=======
$( "#formVK" ).submit(function(event) {
  event.preventDefault();
  
  var $form = $( this ),
	vkidI = $form.find( "input[name='vkid']" ).val(),
	urlI = $form.attr( "action" );
	
  var posting = $.post( urlI, { vklink: vkidI })
  .always(function() {
	$("#loadVK").show();
	$(".vkDAE").attr("disabled");
  })
  .done(function() {
    $("#loadVK").hide();
	$(".vkDAE").removeAttr("disabled");
  })
  .fail(function() {
    $("#loadVK").hide();
	$(".vkDAE").removeAttr("disabled");
  });
>>>>>>> ac09b82467015a144e1e099f0ca6c1b1a1d02bb9
});