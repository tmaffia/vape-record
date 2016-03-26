$(document).ready(function() {
	$('.record-vape').click(function(){
		$.get("api/vape", function() {
			location.reload();
			});
	});
});
