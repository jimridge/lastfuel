#!/usr/bin/env python3
from flask import Flask
from flask import render_template
from flask import jsonify

from gpx_converter import Converter

app = Flask(__name__)
app.config['TEMPLATES_AUTO_RELOAD'] = True

@app.errorhandler(404)
def page_not_found(error):
    return render_template('404.html'), 404


@app.route("/")
@app.route('/<name>')
def index(name=None, title="LastFuel :: Canyoning", year=2020):
    return render_template('index.html', name=name, title=title, year=year)


@app.route("/data/<name>")
def data(name=None):

    if name == "map.json":

        d={"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":["150.3079828","-33.5449968"]},"properties":{"id":"1","name":"Dalpura Canyon", "description": "Dalpura is a relatively short canyon on the north side of the Grose Valley. It has one short abseil and the burnt out sections at the beginning of the canyon are a stark reminder of the late 2019/early 2020 bushfires. The walk in is not particularly tough, but a moderate level of fitness is required for the uphill walk out. A half day canyon!", "status":"Easy"}}]
        }
        return jsonify(d)

    if name == "siloam.json":

        gpx_dict = Converter(input_file='./data/activity_5655491821.gpx').gpx_to_dictionary(latitude_key='latitude', longitude_key='longitude')

        coordinates = []
        for i in range(0, len(gpx_dict['latitude'])):
            # print("{} -- {}".format(gpx_dict['latitude'][i], gpx_dict['longitude'][i]))
            point = [gpx_dict['longitude'][i], gpx_dict['latitude'][i]]
            coordinates.append(point)

        d={ 'type': 'Feature',
            'properties': {"id":"1","name":"Pool of Siloam", "description": "A quick descent to a beautiful small, sheltered pool and waterfall. The water is icy cold, and there is a good climb out to the east to test your fitness!", "status":"Easy"},
            'geometry': {
            'type': 'LineString',
            'coordinates': coordinates,
            }
        }

        return jsonify(d)

    return



# @app.route('/')
# def hello_world():
#     return 'Hello, World!'
# @app.route('/charts/line.svg')
# def line_route():
#     chart = pygal.Line()
#     chart.add('line', [.0002, .0005, .00035])
#     return chart.render_response()

# @app.route('/path/<path:subpath>')
# def show_subpath(subpath):
#     # show the subpath after /path/
#     return '<strong>Subpath:</strong> %s' % subpath