console.log("loaded")

$(window).on('load', function () {
  $('#load').delay(17500).fadeOut('slow');
  $('#block_two').delay(17600).show(0);

  $('#example3').typeIt({
    strings: ["Добро пожаловать на Бета-Тест сайта сервера по майкнрафту.", "SWORD ART ONLINE", "После прогрузки сайта,авторизируйтесь или зарегистрируйтесь", "затем, нажмите кнопку Home, что бы вызвать панель.", "Если у вас ноутбук,зажмите клавишу *Fn и стрелку в лево*"],
    speed: 50,
    autoStart: false
  });
});

