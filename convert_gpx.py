from gpx_converter import Converter

gpx_dict = Converter(input_file='./data/activity_5655491821.gpx').gpx_to_dictionary(latitude_key='latitude', longitude_key='longitude')
# now you have a dictionary and can access the longitudes and latitudes values from the keys
latitudes = gpx_dict['latitude']
longitudes = gpx_dict['longitude']

coordinates = []
for i in range(0, len(gpx_dict['latitude'])):
    # print("{} -- {}".format(gpx_dict['latitude'][i], gpx_dict['longitude'][i]))
    point = [gpx_dict['longitude'][i], gpx_dict['latitude'][i]]
    coordinates.append(point)

print(coordinates)
