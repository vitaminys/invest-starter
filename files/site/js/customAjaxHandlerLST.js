$( "#pay1" ).submit(function(event) {
  event.preventDefault();
  
  var $form = $( this ),
	idI = 1,
	urlI = $form.attr( "action" );
	
  var posting = $.post( urlI, { vklink: idI })
  .always(function() {
	
  })
  .done(function() {
    
  })
  .fail(function() {
    
  });
});