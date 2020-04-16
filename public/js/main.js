(function($){
    $(function(){
          
        // Settings
        var host = window.location.hostname;
    
    
        var ws_url = 'ws://localhost:8080/echo/';
        var connection;
        var ws_waiting = 0;
    
        // ******************************************************************
        // Side navigation
        // ******************************************************************
        $('.button-collapse').sideNav();
        
        // Navlinks
        $('#mc-nav').on('click', '.mc-navlink', function(){
            console.log("Navigate to pane: ", $(this).data("pane"));
            showPane($(this).data("pane"));
        });
        
        function showPane(pane) {
            $('.mc_pane').addClass('hide');
            $('#' + pane).removeClass('hide');
            $('.button-collapse').sideNav('hide');
            
            if (pane == "pane2") {
                setMainColor();
            }
        }
        
        
        // ******************************************************************
        // init()
        // ******************************************************************
        function init() {
            console.log("Connection websockets to:", ws_url);
            connection = new WebSocket(ws_url);
            
        }	
        // ******************************************************************
        // WebSocket commands
        // ******************************************************************
        function wsSendCommand(cmd) {
            console.log("Send WebSocket command:", cmd);
            if (ws_waiting == 0)  {
                connection.send(cmd);
                ws_waiting++;
            } else {
                console.log("++++++++ WS call waiting, skip")
            }
        }	
            
        //handle the touch event
        function doTouch(event) {
            //to not also fire on click
            event.preventDefault();
            var el = event.target;
            
            //touch position
            var pos = {x: Math.round(event.targetTouches[0].pageX - el.offsetLeft),
                       y: Math.round(event.targetTouches[0].pageY - el.offsetTop)};
            //color
            var color = context.getImageData(pos.x, pos.y, 1, 1).data;
    
            updateStatus(pos, color);
        }
    
        function doClick(event) {   
            //click position   
            var pos = getMousePos(canvas, event);
            //color
            var color = context.getImageData(pos.x, pos.y, 1, 1).data;
            
            //console.log("click", pos.x, pos.y, color);
            updateStatus(pos, color);
            
            //now do sth with the color rgbToHex(color);
            //don't do stuff when #000000 (outside circle and lines
        }
    
    
        // ******************************************************************
        // main
        // ******************************************************************
        init();
        
    }); // end of document ready
    })(jQuery); 