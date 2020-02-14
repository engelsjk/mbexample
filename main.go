package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

var CLI struct {
	// Styles (10)
	Map              struct{} `cmd help:"Initialize a map in an HTML element with Mapbox GL JS."`
	ChangeStyle      struct{} `cmd help:"Switch to another map style."`
	SatelliteMap     struct{} `cmd help:"Satellite raster baselayer."`
	WorldCopies      struct{} `cmd help:"Toggle between rendering a single world and multiple copies of the world."`
	CustomStyle      struct{} `cmd help:"Using a custom Mapbox-hosted style."`
	AddIcon          struct{} `cmd help:"Add an icon to the map from an external URL and use it in a symbol layer."`
	StretchableImage struct{} `cmd help:"Use a stretchable image as a background for text."`
	MissingIcon      struct{} `cmd help:"Dynamically generate a missing icon at runtime and add it to the map."`
	GeneratedIcon    struct{} `cmd help:"Add an icon to the map that was generated at runtime."`
	AnimatedIcon     struct{} `cmd help:"Add an animated icon to the map that was generated at runtime with a Canvas."`
	// Layers (32)
	BuildingsThreedee  struct{} `cmd help:"Use extrusions to display buildings' height in 3D."`
	IndoorThreedee     struct{} `cmd help:"Create a 3D indoor map with the fill-extrude-height paint property."`
	ModelThreedee      struct{} `cmd help:"Use a custom style layer with three.js to add a 3D model to the map."`
	LayerOpacity       struct{} `cmd help:"Drag the range slider to adjust the opacity of a raster layer on top of a map."`
	AnimateLine        struct{} `cmd help:"Animate a line by updating a GeoJSON source on each frame."`
	AnimateImages      struct{} `cmd help:"Use a series of image sources to create an animation."`
	AnimatePoint       struct{} `cmd help:"Animate the position of a point by updating a GeoJSON source on each frame."`
	BuildingColor      struct{} `cmd help:"Uses the interpolate expression to ease-in the building layer and smoothly fade from one color to the next."`
	LabelsCase         struct{} `cmd help:"Use the upcase and downcase expressions to change the case of labels."`
	HtmlClusters       struct{} `cmd help:"Uses Mapbox GL JS clustering with HTML markers and custom property expressions."`
	CreateClusters     struct{} `cmd help:"Use Mapbox GL JS' built-in functions to visualize points as clusters."`
	LayerColor         struct{} `cmd help:"Using setPaintProperty to change a layer's fill color."`
	CustomWebgl        struct{} `cmd help:"Use a custom style layer to render custom WebGL content."`
	StyleCircles       struct{} `cmd help:"Creating a visualization with a data expression for circle-color."`
	StyleLines         struct{} `cmd help:"Creating a visualization with a data expression for line-color."`
	StyleLabels        struct{} `cmd help:"Uses the format expression to display country labels in both English and in the local language."`
	FallbackImage      struct{} `cmd help:"Fallback to another image when a requested image is not available."`
	PolygonPattern     struct{} `cmd help:"Use fill-pattern to draw a polygon from a repeating image pattern."`
	GeojsonLine        struct{} `cmd help:"Add a GeoJSON line to a map."`
	LayerLabels        struct{} `cmd help:"Using the second argument of addLayer, you can be more precise."`
	GeojsonPoints      struct{} `cmd help:"Draw points from a GeoJSON collection to a map."`
	GeojsonPolygon     struct{} `cmd help:"Style a polygon with the fill layer type."`
	HeatmapLayer       struct{} `cmd help:"Visualize earthquake frequency by location using a heatmap layer."`
	AddHillshading     struct{} `cmd help:"Adds raster hillshading to a map."`
	GradientLine       struct{} `cmd help:"Use the line-gradient paint property and an expression to visualize distance from the starting point of a line."`
	MultipleGeometries struct{} `cmd help:"Add a polygon and circle layer from the same GeoJSON source."`
	StyleBathymetry    struct{} `cmd help:"Uses the interpolate expression with a cubic bezier curve expression to style bathymetry data."`
	HideLayers         struct{} `cmd help:"Create a custom layer switcher to display different datasets."`
	ChangeWorldview    struct{} `cmd help:"Uses the worldview value to adjust administrative boundaries based on the map's audience."`
	ZoomChoropleth     struct{} `cmd help:"Using 2014 census data, display state or county population depending on zoom level."`
	PopulationDensity  struct{} `cmd help:"Using a variable binding expression to calculate and display population density."`
	LabelPlacement     struct{} `cmd help:"Use text-variable-anchor to allow high priority labels to shift position to stay on the map."`
	// Sources (9)
	MergeGeometries struct{} `cmd help:"Style a choropleth by merging local JSON data with vector tile geometries."`
	AddImage        struct{} `cmd help:"Dark vector baselayer with radar weather image overlay."`
	RealtimeData    struct{} `cmd help:"Use realtime GeoJSON data streams to move a symbol on your map."`
	RealtimeFeature struct{} `cmd help:"Change an existing feature on your map in realtime by updating its data."`
	AddRaster       struct{} `cmd help:"Add a third-party raster source to the map."`
	OutsideVector   struct{} `cmd help:"Render vector data provided by Mapillary."`
	AddVideo        struct{} `cmd help:"Satellite raster baselayer with video on top. Click on the map to play and pause."`
	AddWms          struct{} `cmd help:"Adding an external Web Map Service layer to the map."`
	AddVector       struct{} `cmd help:"Add a vector source to a map."`
	// User Interaction (20)
	BuildingRhythm     struct{} `cmd help:"A map where the 3D buildings dance to the rhythm of your ambient environment."`
	DisableRotation    struct{} `cmd help:"Prevent users from rotating a map."`
	DraggableMarker    struct{} `cmd help:"Drag the Marker to a new location on a map and populates its coordinates in a display."`
	DraggablePoint     struct{} `cmd help:"Drag the point to a new location on a map and populates its coordinates in a display."`
	FilterFeatures     struct{} `cmd help:"Move the map to query viewable features in a vector tile layer and filter by typing in an input."`
	FilterSymbols      struct{} `cmd help:"Filter symbols by icon name by typing in a text input."`
	ToggleSymbols      struct{} `cmd help:"Filter a set of symbols based on a property value in the data."`
	HoverEffect        struct{} `cmd help:"Using events and feature states to create a per feature hover effect."`
	StaticMap          struct{} `cmd help:"Setting interactive: false to create a static map."`
	SwitchLanguage     struct{} `cmd help:"Using setLayoutProperty to switch languages dynamically."`
	MeasureDistance    struct{} `cmd help:"Click points on a map to create lines that measure distanced using turf.length."`
	MousePosition      struct{} `cmd help:"Showing mouse position on hover with pixel and latitude and longitude coordinates."`
	HighlightFeatures  struct{} `cmd help:"Hover over counties to highlight counties that share the same name."`
	QueryFeatures      struct{} `cmd help:"Click on the map to query features using queryRenderedFeatures."`
	HoverQuery         struct{} `cmd help:"Using queryRenderedFeatures to show properties of hovered-over map elements."`
	RestrictPanning    struct{} `cmd help:"Prevent a map from being panned to a different place by setting maxBounds."`
	TimeSlider         struct{} `cmd help:"Drag the range slider to visualize earthquakes that were greater than 5.9 magnitude in 2015."`
	ToggleInteractions struct{} `cmd help:"Enable or disable UI handlers on a map."`
	BoxQuery           struct{} `cmd help:"Hold the Shift key and drag the map to query features using queryRenderedFeatures."`
	FitLinestring      struct{} `cmd help:"Get the bounds of a LineString by passing its first coordinates to LngLatBounds and chaining extend to include the last coordinates."`
	// Camera (11)
	AnimateCamera struct{} `cmd help:"Animate map camera around a point."`
	AnimateRoute  struct{} `cmd help:"Use Turf to smoothly animate a point along the distance of a line."`
	CenterMap     struct{} `cmd help:"Using events and flyTo to center the map on a symbol."`
	FitBounds     struct{} `cmd help:"Use fitBounds to show a specific area of the map in view, regardless of the pixel size of the map."`
	FlySlow       struct{} `cmd help:"Using flyTo with flyOptions."`
	FlyTo         struct{} `cmd help:"Using flyTo to smoothly interpolate between locations."`
	JumpTo        struct{} `cmd help:"Use the jumpTo function to showcase multiple locations."`
	MapSlideshow  struct{} `cmd help:"Autoplay the locations of boroughs in New York City."`
	FlyScroll     struct{} `cmd help:"Scroll down through the story and the map will fly to the chapter's location."`
	PitchBearing  struct{} `cmd help:"Map options extend CameraOptions, so you can set more than just center and zoom."`
	MapGame       struct{} `cmd help:"Move around the map with game-like controls."`
	// Controls and Overlays (17)
	AnimateMarker      struct{} `cmd help:"Animate the position of a Marker by updating its location on each frame."`
	AttributePosition  struct{} `cmd help:"Place attribution in the top-left position when initializing a map."`
	CustomIcons        struct{} `cmd help:"Use Marker to add custom icons to your map."`
	DisableZoom        struct{} `cmd help:"Prevents scroll from zooming a map."`
	Fullscreen         struct{} `cmd help:"Toggle between current view and fullscreen mode."`
	LocateUser         struct{} `cmd help:"Geolocate the user and then track their current location on the map using the GeolocateControl."`
	CompareMaps        struct{} `cmd help:"Use mapbox-gl-compare to swipe between and synchronize two maps."`
	DrivingDirections  struct{} `cmd help:"Use the mapbox-gl-directions plugin to show results from the Mapbox Directions API. Click the map to add an origin and destination."`
	DrawPolygon        struct{} `cmd help:"Use mapbox-gl-draw to draw a polygon and Turf.js to calculate its area in square meters."`
	AddGeocoder        struct{} `cmd help:"Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API."`
	MarkGeocode        struct{} `cmd help:"Add a Marker using a place name or address for its location using the forward geocoder."`
	NavigationControls struct{} `cmd help:"Zoom and rotation controls to make map navigation more obvious."`
	ShowPopup          struct{} `cmd help:"When a user clicks a polygon, show a Popup containing more information."`
	SymbolPopup        struct{} `cmd help:"When a user clicks a symbol, show a Popup containing more information."`
	HoverPopup         struct{} `cmd help:"When a user hovers over a symbol, show a Popup containing more information."`
	AddPopup           struct{} `cmd help:"Add a Popup to the map."`
	MarkerPopup        struct{} `cmd help:"Attach a Popup to a Marker and display it on click."`

	// Geocoder (8)
	GeocoderCoordinates struct{} `cmd help:"Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API, while accepting geographic coordinates provided as input."`
	GeocoderDropdown    struct{} `cmd help:"Use a custom html rendering function with the mapbox-gl-geocoder to customize how the dropdown menu is displayed."`
	GeocoderRegion      struct{} `cmd help:"Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API limiting results to a region."`
	GeocoderOutside     struct{} `cmd help:"Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API attached to an element outside the map."`
	GeocoderCamera      struct{} `cmd help:"Use camera animation options with the mapbox-gl-geocoder to create a custom animation when a result is selected."`
	GeocoderLanguage    struct{} `cmd help:"Localize the mapbox-gl-geocoder to set the UI language and improve result relevance in that language."`
	GeocoderMarker      struct{} `cmd help:"Style a Marker used to mark the location of the mapbox-gl-geocoder result."`
	GeocoderSupplement  struct{} `cmd help:"Use the mapbox-gl-geocoder control to search for places using Mapbox Geocoding API, supplementing results from a local data source or function."`
	// Browser Support (1)
	BrowserSupport struct{} `cmd help:"Check for Mapbox GL browser support."`
	// Internationalization Support (2)
	LocalFonts struct{} `cmd help:"Use the localIdeographFontFamily setting to speed up map load times by using locally available fonts instead of font data fetched from the server."`
	RightLeft  struct{} `cmd help:"Use the mapbox-gl-rtl-text plugin to support the Arabic or Hebrew languages, which are written right-to-left. "`
	// Flags
	Token string `help:"Mapbox access token."`
	Save  bool   `help:"Save map html file locally."`
	Style string `help:"Specify Mapbox standard style." enum:",streets,light,dark,outdoors,satellite,satellite-streets"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("mbexample"),
		kong.Description("Quick and easy Mapbox GL JS maps example templates."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	GetExample(ctx.Command())
}

func GetExample(example string) {
	m, err := NewMap(example, CLI.Style)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if CLI.Save {
		err = m.Save()
		if err != nil {
			fmt.Printf("warning: example map file %s.html could not be saved\n", m.Name)
		} else {
			fmt.Printf("example %s.html saved!\n", m.Name)
		}
	}
	http.Handle("/", m)
	port := 3000
	fmt.Printf("example %s at localhost:%d...\n", m.Name, port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	errPortInUse := fmt.Sprintf("listen tcp :%d: bind: address already in use", port)
	if err.Error() == errPortInUse {
		fmt.Printf("error: nevermind, port %d seems to be in use already\n", port)
		return
	}
}

func NewMap(example string, style string) (*Map, error) {
	accessToken, err := getAccessToken()
	if err != nil {
		return nil, err
	}
	templateFilename := example + ".gohtml"
	t, err := vfstemplate.ParseFiles(assets, nil, templateFilename)
	if err != nil {
		return nil, fmt.Errorf("unable to find a template for the %s map", example)
	}
	return &Map{
		Name:             example,
		AccessToken:      accessToken,
		Style:            getStyleID(style),
		Template:         t,
		TemplateFileName: templateFilename,
	}, nil
}

type Map struct {
	Name             string
	AccessToken      string
	Style            string
	Template         *template.Template
	TemplateFileName string
}

func (m *Map) Save() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(path, m.Name+".html"))
	if err != nil {
		return err
	}
	defer f.Close()

	err = m.Template.Execute(f, m)
	if err != nil {
		return (err)
	}
	return nil
}

func (m *Map) Render(w http.ResponseWriter) error {
	return m.Template.ExecuteTemplate(w, m.TemplateFileName, m)
}

func (m *Map) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Render(w); err != nil {
		panic(err)
	}
}

func getAccessToken() (string, error) {
	token := CLI.Token
	if token == "" {
		token = os.Getenv("MAPBOX_ACCESS_TOKEN")
	}
	if token == "" {
		return "", fmt.Errorf("mapbox access token required (use either a MAPBOX_ACCESS_TOKEN env variable or the --token flag)")
	}
	return token, nil
}

func getStyleID(style string) string {
	switch style {
	case "streets":
		return "streets-v11"
	case "light":
		return "light-v10"
	case "dark":
		return "dark-v10"
	case "outdoors":
		return "outdoors-v11"
	case "satellite":
		return "satellite-v9"
	case "satellite-streets":
		return "satellite-streets-v11"
	default:
		return "streets-v11"
	}
}
