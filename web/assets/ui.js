(function() {

  var DOT_RADIUS = 0.03;
  var SVG_PRECISION = 3;
  var SVG_NAMESPACE = 'http://www.w3.org/2000/svg';

  var points = [];
  var curveElement, pointsElement;
  var pointElements = [];
  var curveLabel;

  function main() {
    var reset = document.getElementById('reset-button');
    reset.addEventListener('click', resetPoints);

    curveLabel = document.getElementById('curve-code');

    curveElement = document.getElementById('curve');
    pointsElement = document.getElementById('points');

    var preview = document.getElementById('preview');
    preview.addEventListener('click', function(e) {
      var bounding = preview.getBoundingClientRect();
      var posX = (e.clientX - bounding.left) / bounding.width;
      var posY = (e.clientY - bounding.top) / bounding.height;
      addPoint(posX, 1-posY);
    });

    generateCurve();
  }

  function addPoint(x, y) {
    var dot = document.createElementNS(SVG_NAMESPACE, 'circle');
    dot.setAttribute('cx', x.toFixed(SVG_PRECISION));
    dot.setAttribute('cy', (1-y).toFixed(SVG_PRECISION));
    dot.setAttribute('r', DOT_RADIUS.toFixed(SVG_PRECISION));
    points.push({x: x, y: y});
    pointElements.push(dot);
    pointsElement.appendChild(dot);
    generateCurve();
  }

  function resetPoints() {
    points = [];
    for (var i = 0, len = pointElements.length; i < len; ++i) {
      pointsElement.removeChild(pointElements[i]);
    }
    pointElements = [];
    generateCurve();
  }

  function generateCurve() {
    if (points.length === 0) {
      curveElement.setAttribute('d', 'M0,1 L1,0');
      curveLabel.textContent = 'cubic-bezier(0,0,0,0)';
      return;
    }

    curveElement.setAttribute('d', '');

    window.fitBezierCurve(points, function(curve) {
      var bezierCode = curve[0].toFixed(SVG_PRECISION) + ',' +
        (1-curve[1]).toFixed(SVG_PRECISION) + ',' +
        curve[2].toFixed(SVG_PRECISION) + ',' +
        (1-curve[3]).toFixed(SVG_PRECISION);

      var curveData = 'M0,1 C' + bezierCode + ' 1,0';
      curveElement.setAttribute('d', curveData);

      curveLabel.textContent = 'cubic-bezier(' + bezierCode + ')';
    });
  }

  window.addEventListener('load', main);

})();
