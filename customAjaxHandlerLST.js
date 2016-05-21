<<<<<<< HEAD
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
=======
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
>>>>>>> ac09b82467015a144e1e099f0ca6c1b1a1d02bb9
});