$ ->
  $('#contestant_name').focus()
  $('#contestants').on 'hover', '.contestant:has(.retire)', (event) ->
    if event.type == 'mouseenter'
      $(this).find('.retire').show()
    else
      $(this).find('.retire').hide()
  $('#new-contestant-form').submit (event) ->
    inputField = $('#contestant_name')
    if inputField.val().length == 0
      event.preventDefault()
      inputField.focus()
      return false
