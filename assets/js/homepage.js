((app, $) => {
  $(() => {
    app();
  });
})(() => {
  'use strict';

  var $container = $('div.container');

  // Define toolbar block component
  var toolbarBlock = {
    el: {
      body: '.cr-roomlist-toolbar',
      modal: '#createModalBlock',
    },
    /**
     * Bind event listener
     * 
     * @param {object} $el
     * @param {object} $modal
     */
    bindEvents: ($el, $modal) => {

      // new a chatroom
      $modal.find('.newroom-check-btn').on('click', () => {
        $.ajax({
          type: 'POST',
          url: 'chatrooms',
          dataType: 'json',
          data: $modal.find('form').serialize(),
        }).done((res) => {
          console.log('success', res)
          // reload the table
          roomListBlock.fetchData();
          // close the modal
          $modal.find('.newroom-close-btn').trigger('click');
        }).fail((xhr) => {
          console.log('fail', xhr.responseJSON, xhr.responseText)
        });
      });
    },
    /**
     * Initialize
     */
    init: () => {
      let $el = $(toolbarBlock.el.body);
      let $modal = $(toolbarBlock.el.modal);
      toolbarBlock.bindEvents($el, $modal);
    },
  };

  // Define room list block component
  var roomListBlock = {
    $el: $('#cr-roomlist-table'),
    options: {
      table: {
        info: true,
        paging: true,
        searching: true,
        ordering: true,
        autoWidth: false,
        deferRender: true,
        destroy: true,
        columns: [
          {
            // chatroom name
            className: 'text-center',
            data: 'chatroomName',
          },
          {
            // room chief
            className: 'text-center',
            data: 'CreateBy',
          },
          {
            // operation
            className: 'text-center',
            data: null,
            render: (data, type, row, meta) => {
              return $('<a>')
                .addClass('link-info room-entry')
                .text('進入')
                .prop('outerHTML');
            },
          }
        ],
        createdRow: (row, data, dataIndex, cells) => {
          let $row = $(row);

          // attach click event to entry the chatroom
          $row.find('.room-entry').on('click', () => {

            // entry to the chatroom
            $('<form>')
              .attr('method', 'POST')
              .attr('action', '/chatroom')
              .attr('target', '_blank')
              .append($('<input>')
                .attr({
                  type: 'hidden',
                  name: 'RoomID',
                  value: data.ID,
                })
              )
              .appendTo('body')
              .submit();
          });
        },
      }
    },
    /**
     * Fetch the info of active chatrooms
     */
    fetchData: () => {
      $.ajax({
        type: 'GET',
        url: '/chatrooms',
        dataType: 'json',
      }).done((res) => {
        let data = res.data || {};
        let tableOptions = $.extend(
          {},
          roomListBlock.options.table, {
          data: data,
        });
        roomListBlock.renderTable(tableOptions);
      }).fail((xhr) => {
        console.log('fail', xhr.responseJSON, xhr.responseText, xhr)
      })
    },
    /**
     * Render the datatable
     * 
     * @param {object} option
     */
    renderTable: (option) => {
      let $table = roomListBlock.$el;

      if ($.fn.DataTable.isDataTable($table)) {
        $table.DataTable().destroy();
      }

      roomListBlock.$el.DataTable(option);
    },
    /**
     * Initialize
     */
    init: () => {
      roomListBlock.fetchData();
    },
  };

  // test ws 
  // TODO
  $('.cr-link').on('click', () => {
    var loc = window.location;
    var uri = loc.protocol === 'https:' ? 'wss:' : 'ws:';

    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    console.log('loc', loc);
    console.log('uri', uri);

    let ws = new WebSocket(uri)

    ws.onopen = () => {
      console.log('Websocket Connected!!')
    }

    ws.onmessage = (evt) => {
      console.log('ws return data: ', evt)
      var out = document.getElementById('output');
      out.innerHTML += evt.data + '<br>';
    }

    setInterval(() => {
      ws.send('Hello, Server!');
    }, 1000);
  });

  // Constructor
  var constructor = () => {

    // activate toolbar
    toolbarBlock.init();

    // activate room list
    roomListBlock.init();
  };

  constructor();
}, jQuery)