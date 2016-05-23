(function() {

  var fitID = 0;
  var worker = new Worker('assets/webworker.js');
  var callbacks = {};

  worker.onmessage = function(e) {
    var handler = callbacks[e.data[0]];
    delete callbacks[e.data[0]];
    handler(e.data.slice(1));
  };

  function fitBezierCurve(points, callback) {
    callbacks[fitID] = callback;
    worker.postMessage([fitID].concat(points));
    ++fitID;
  }

  window.fitBezierCurve = fitBezierCurve;

})();
