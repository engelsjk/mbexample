# mbexample

Quick and easy Mapbox GL JS map templates.

```bash
go get -u github.com/engelsjk/mbexample
export MAPBOX_ACCESS_TOKEN='***ACCESSTOKEN***'
mbexample map
```

# Why?

Find yourself frequently copying [Mapbox examples](https://docs.mapbox.com/mapbox-gl-js/examples/) to start new map ideas? Then why not use a convenient if not unnecessary CLI to spin up local copies of any Mapbox GL JS example!

# How?

A few optional flags can be used to save an html copy locally, set a Mapbox access token or to specify a standard Mapbox styles.

```
mbexample [EXAMPLE] --save --token=***ACCESSTOKEN*** --style=[streets, light, dark, outdoors, satellite or satellite-streets]
```

The example map you specified should then be viewable at localhost:3000.


# Examples

Here's a full list of examples you can choose from, copied directly from the official list of [Mapbox GL JS Examples](https://docs.mapbox.com/mapbox-gl-js/examples/).

## Styles
```
map                     Initialize a map in an HTML element with Mapbox GL JS.
change-style            Switch to another map style.
satellite-map           Satellite raster baselayer.
world-copies            Toggle between rendering a single world and multiple copies of the world.
custom-style            Using a custom Mapbox-hosted style.
add-icon                Add an icon to the map from an external URL and use it in a symbol layer.
stretchable-image       Use a stretchable image as a background for text.
missing-icon            Dynamically generate a missing icon at runtime and add it to the map.
generated-icon          Add an icon to the map that was generated at runtime.
animated-icon           Add an animated icon to the map that was generated at runtime with a Canvas.
```

## Layers
```
buildings-threedee      Use extrusions to display buildings' height in 3D.
indoor-threedee         Create a 3D indoor map with the fill-extrude-height paint property.
model-threedee          Use a custom style layer with three.js to add a 3D model to the map.
layer-opacity           Drag the range slider to adjust the opacity of a raster layer on top of a map.
animate-line            Animate a line by updating a GeoJSON source on each frame.
animate-images          Use a series of image sources to create an animation.
animate-point           Animate the position of a point by updating a GeoJSON source on each frame.
building-color          Uses the interpolate expression to ease-in the building layer and smoothly fade from one color to the next.
labels-case             Use the upcase and downcase expressions to change the case of labels.
html-clusters           Uses Mapbox GL JS clustering with HTML markers and custom property expressions.
create-clusters         Use Mapbox GL JS' built-in functions to visualize points as clusters.
layer-color             Using setPaintProperty to change a layer's fill color.
custom-webgl            Use a custom style layer to render custom WebGL content.
style-circles           Creating a visualization with a data expression for circle-color.
style-lines             Creating a visualization with a data expression for line-color.
style-labels            Uses the format expression to display country labels in both English and in the local language.
fallback-image          Fallback to another image when a requested image is not available.
polygon-pattern         Use fill-pattern to draw a polygon from a repeating image pattern.
geojson-line            Add a GeoJSON line to a map.
layer-labels            Using the second argument of addLayer, you can be more precise.
geojson-points          Draw points from a GeoJSON collection to a map.
geojson-polygon         Style a polygon with the fill layer type.
heatmap-layer           Visualize earthquake frequency by location using a heatmap layer.
add-hillshading         Adds raster hillshading to a map.
gradient-line           Use the line-gradient paint property and an expression to visualize distance from the starting point of a line.
multiple-geometries     Add a polygon and circle layer from the same GeoJSON source.
style-bathymetry        Uses the interpolate expression with a cubic bezier curve expression to style bathymetry data.
hide-layers             Create a custom layer switcher to display different datasets.
change-worldview        Uses the worldview value to adjust administrative boundaries based on the map's audience.
zoom-choropleth         Using 2014 census data, display state or county population depending on zoom level.
population-density      Using a variable binding expression to calculate and display population density.
label-placement         Use text-variable-anchor to allow high priority labels to shift position to stay on the map.
 ``` 
 
## Sources
```
merge-geometries        Style a choropleth by merging local JSON data with vector tile geometries.
add-image               Dark vector baselayer with radar weather image overlay.
realtime-data           Use realtime GeoJSON data streams to move a symbol on your map.
realtime-feature        Change an existing feature on your map in realtime by updating its data.
add-raster              Add a third-party raster source to the map.
outside-vector          Render vector data provided by Mapillary.
add-video               Satellite raster baselayer with video on top. Click on the map to play and pause.
add-wms                 Adding an external Web Map Service layer to the map.
add-vector              Add a vector source to a map.
```

## User Interaction
```
building-rhythm         A map where the 3D buildings dance to the rhythm of your ambient environment.
disable-rotation        Prevent users from rotating a map.
draggable-marker        Drag the Marker to a new location on a map and populates its coordinates in a display.
draggable-point         Drag the point to a new location on a map and populates its coordinates in a display.
filter-features         Move the map to query viewable features in a vector tile layer and filter by typing in an input.
filter-symbols          Filter symbols by icon name by typing in a text input.
toggle-symbols          Filter a set of symbols based on a property value in the data.
hover-effect            Using events and feature states to create a per feature hover effect.
static-map              Setting interactive: false to create a static map.
switch-language         Using setLayoutProperty to switch languages dynamically.
measure-distance        Click points on a map to create lines that measure distanced using turf.length.
mouse-position          Showing mouse position on hover with pixel and latitude and longitude coordinates.
highlight-features      Hover over counties to highlight counties that share the same name.
query-features          Click on the map to query features using queryRenderedFeatures.
hover-query             Using queryRenderedFeatures to show properties of hovered-over map elements.
restrict-panning        Prevent a map from being panned to a different place by setting maxBounds.
time-slider             Drag the range slider to visualize earthquakes that were greater than 5.9 magnitude in 2015.
toggle-interactions     Enable or disable UI handlers on a map.
box-query               Hold the Shift key and drag the map to query features using queryRenderedFeatures.
fit-linestring          Get the bounds of a LineString by passing its first coordinates to LngLatBounds and chaining extend            to include the last coordinates.
```

## Camera
```
animate-camera          Animate map camera around a point.
animate-route           Use Turf to smoothly animate a point along the distance of a line.
center-map              Using events and flyTo to center the map on a symbol.
fit-bounds              Use fitBounds to show a specific area of the map in view, regardless of the pixel size of the map.
fly-slow                Using flyTo with flyOptions.
fly-to                  Using flyTo to smoothly interpolate between locations.
jump-to                 Use the jumpTo function to showcase multiple locations.
map-slideshow           Autoplay the locations of boroughs in New York City.
fly-scroll              Scroll down through the story and the map will fly to the chapter's location.
pitch-bearing           Map options extend CameraOptions, so you can set more than just center and zoom.
map-game                Move around the map with game-like controls.
```

## Controls and Overlays
```
animate-marker          Animate the position of a Marker by updating its location on each frame.
attribute-position      Place attribution in the top-left position when initializing a map.
custom-icons            Use Marker to add custom icons to your map.
disable-zoom            Prevents scroll from zooming a map.
fullscreen              Toggle between current view and fullscreen mode.
locate-user             Geolocate the user and then track their current location on the map using the GeolocateControl.
compare-maps            Use mapbox-gl-compare to swipe between and synchronize two maps.
driving-directions      Use the mapbox-gl-directions plugin to show results from the Mapbox Directions API. Click the map to add an origin and destination.
draw-polygon            Use mapbox-gl-draw to draw a polygon and Turf.js to calculate its area in square meters.
add-geocoder            Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API.
mark-geocode            Add a Marker using a place name or address for its location using the forward geocoder.
navigation-controls     Zoom and rotation controls to make map navigation more obvious.
show-popup              When a user clicks a polygon, show a Popup containing more information.
symbol-popup            When a user clicks a symbol, show a Popup containing more information.
hover-popup             When a user hovers over a symbol, show a Popup containing more information.
add-popup               Add a Popup to the map.
marker-popup            Attach a Popup to a Marker and display it on click.
```

## Geocoder
```
geocoder-coordinates    Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API, while accepting geographic coordinates provided as input.
geocoder-dropdown       Use a custom html rendering function with the mapbox-gl-geocoder to customize how the dropdown menu is displayed.
geocoder-region         Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API limiting results to a region.
geocoder-outside        Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API attached to an element outside the map.
geocoder-camera         Use camera animation options with the mapbox-gl-geocoder to create a custom animation when a result is selected.
geocoder-language       Localize the mapbox-gl-geocoder to set the UI language and improve result relevance in that language.
geocoder-marker         Style a Marker used to mark the location of the mapbox-gl-geocoder result.
geocoder-supplement     Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API, supplementing results from a local data source or function.
```

## Browser Support 
```
browser-support         Check for Mapbox GL browser support.
```

## Internationalization Support
```
local-fonts             Use the localIdeographFontFamily setting to speed up map load times by using locally available fonts instead of font data fetched from the server.
right-left              Use the mapbox-gl-rtl-text plugin to support the Arabic or Hebrew languages, which are written right-to-left.
```
