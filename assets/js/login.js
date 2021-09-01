((app, $) => {
  $(() => {
    app();
  });
})(() => {
  'use strict';

  var $container = $('div.container');

  // Define the register modal component
  var registerModal = {
    el: '#registerModalBlock',
    /**
     * Form validator
     * 
     * @param   {object}  $el 
     * @returns {boolean}
     */
    validator: ($el) => {
      let result = true;
      let columns = {
        $account: $el.find('input[name="account"]'),
        $pwd: $el.find('input[name="password"]'),
        $pwd2: $el.find('input[name="password2"]'),
        $nickname: $el.find('input[name="nickname"]'),
        $email: $el.find('input[name="email"]'),
      };

      // Check all required inputs have been filled in
      $.each(columns, (key, $col) => {
        let $label = $col.prev();
        if ($col.val() === '') {
          // clean old status
          $label.find('span').remove();
          // filed in error message
          $label.append($('<span>').addClass('cr-error-msg').text(' * 必填'));
          // alter validate result
          result = false;
        } else {
          // clean old status
          $label.find('span').remove();
        }
      });

      // Check that the password is consistent with the secondary password
      let pwdMatched = columns.$pwd.val() === columns.$pwd2.val();
      if (!pwdMatched) {
        let $label = columns.$pwd.prev();
        // clean old status
        $label.find('span').remove();
        // filed in error message
        $label.append($('<span>').addClass('cr-error-msg').text(' * 密碼不一致'));
        // alter validate result
        result = false;
      }

      return result;
    },
    /**
     * Bind event listener
     * 
     * @param {object} $el 
     */
    bindEvents: ($el) => {

      // confirm button of register modal
      $el.find('.register-check-button').on('click', $container, () => {
        if (registerModal.validator($el)) {
          $.ajax({
            type: 'POST',
            url: '/users',
            data: $el.find('form').serialize(),
          }).done((res) => {
            console.log('success', res)
          }).fail((xhr) => {
            console.log('fail', xhr.responseJSON, xhr.responseText)
          });
        }
      });
    },
    /**
     * Initialize
     */
    init: () => {
      let $el = $(registerModal.el);
      registerModal.bindEvents($el);
    },
  };

  // Define the login panel component
  var loginPanel = {
    el: '.cr-login-panel',
    /**
     * Bind event listener
     * 
     * @param {object} $el 
     */
    bindEvents: ($el) => {

      // login button
      $el.find('.login-button').on('click', () => {
        $.ajax({
          type: 'POST',
          url: '/login',
          data: $el.find('form').serialize(),
        }).done((res) => {
          console.log('success', res)
          // login success & redirect to homepage
          window.location.href = '/homepage';
        }).fail((xhr) => {
          console.log('fail', xhr.responseJSON, xhr.responseText)
        });
      });
    },
    /**
     * Initialize
     */
    init: () => {
      let $el = $(loginPanel.el);

      loginPanel.bindEvents($el);
    },
  };

  // Constructor
  var constructor = () => {

    // activate login panel
    loginPanel.init();

    // activate register modal
    registerModal.init();
  };

  constructor();

}, jQuery)