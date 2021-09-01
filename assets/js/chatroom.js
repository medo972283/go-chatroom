((app, $) => {
  $(() => {
    app();
  });
})(() => {
  'use strict';

  var $container = $('div.container');

  // Define room title banner component
  var roomTitleBanner = {
    $el: $('.room-title'),
    render: (data) => {
      // chatroom name
      roomTitleBanner.$el
        .find('.chatroom-name')
        .append($('<span>')
          .text(data.chatroom.chatroomName)
        );

      // chatroom host's name
      roomTitleBanner.$el
        .find('.chatroom-host')
        .append($('<span>')
          .text(data.user.nickname)
        );
    },
  };

  // Define room panel component
  var roomPanel = {
    $el: $('.room-panel'),
    // room ID
    roomID: null,
    // web socket instance
    ws: null,
    /**
     * Connect to web socket
     */
    activateWebSocket: () => {
      let loc = window.location;
      let uri = loc.protocol === 'https:' ? 'wss:' : 'ws:';

      uri += '//' + loc.host + '/chatroom/ws';

      if (roomPanel.ws) {
        roomPanel.ws.close();
        roomPanel.ws = null;
      }

      roomPanel.ws = new WebSocket(uri);

      roomPanel.ws.onopen = () => {
        console.log('Websocket Connected!!');
      };

      roomPanel.ws.onclose = () => {
        console.log('Websocket Disconnected!!');
      };

      // render the newest message
      let $messageList = roomPanel.$el.find('#chatroomMsgList');
      roomPanel.ws.onmessage = (evt) => {

        // clear the input message
        msgBlock.$el.find('.send-text').val('');

        // render the latest message element
        roomPanel.renderMsg($messageList, JSON.parse(evt.data));
      };
    },
    /**
     * Render the newest message
     *
     * @param {object} $el
     * @param {object} data
     */
    renderMsg: ($el, data) => {
      let $wrapper = $('<div>').addClass('toast show');
      let $header = $('<div>').addClass('toast-header')
        .append($('<strong>').addClass('me-auto').text(data.UserName))
        .append($('<small>').addClass('text-muted').text(data.CreateAt));
      let $body = $('<div>').addClass('toast-body').text(data.Content);

      $el.append(
        $wrapper
          .append($header)
          .append($body)
      );
    },
    /**
     * Render for the initial reload
     * 
     * @param {object} data 
     */
    initRender: (data) => {

      // render room title
      roomTitleBanner.render(data);

      // render the list of participants in this chatroom
      let $participantList = roomPanel.$el.find('.participants-list');
      $.each(data.participatns, (idx, row) => {
        $participantList
          .append($('<span>')
            .addClass('participant-tag')
            .text(row.Name)
          );
      });

      // render the historical records of the messages
      let $messageList = roomPanel.$el.find('#chatroomMsgList');
      $.each(data.messages, (idx, row) => {
        roomPanel.renderMsg($messageList, row);
      });
    },
    /**
     * Fetch and render the chatroom
     */
    reload: () => {
      $.ajax({
        url: 'chatrooms' + '/' + roomPanel.roomID,
        type: 'GET',
        dataType: 'json',
      }).done((res) => {
        let data = res.data || {};

        // render participants
        roomPanel.initRender(data);

        // activate web socket connection
        roomPanel.activateWebSocket();

      }).fail((xhr) => {
        console.log('fail', xhr.responseJSON, xhr.responseText);
      });
    },
    /**
     * Initialize
     *
     * @param {string} roomID
     */
    init: (roomID) => {
      roomPanel.roomID = roomID;
      roomPanel.reload();
    },
  };

  // Define message sending block
  var msgBlock = {
    $el: $('.send-block'),
    /**
     * Bind event listener
     */
    bindEvents: () => {
      let $el = msgBlock.$el;

      $el.find('button').on('click', () => {

        let $form = $el.find('form');

        // send new message to the web socket server
        roomPanel.ws.send(JSON.stringify({
          content: $form.find('textarea').val(),
          roomID: roomPanel.roomID,
        }));

      });
    },
    /**
     * Initialize
     */
    init: () => {
      msgBlock.bindEvents();
    },
  };

  // Constructor
  var constructor = () => {

    // activate room panel component
    roomPanel.init(ROOM_ID); 

    // activate message block component
    msgBlock.init();
  };

  constructor();
}, jQuery)