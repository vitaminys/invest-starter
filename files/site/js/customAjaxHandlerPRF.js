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
	alert("Fail");
    $("#loadVK").hide();
	$(".vkDAE").removeAttr("disabled");
  });
});
