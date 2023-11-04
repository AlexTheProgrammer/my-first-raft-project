package main

import d "mfrp/dom"

func main() {

	ch := make(chan bool)
	d := d.DOM{
		Body: d.Body(
			d.Div(
				d.Props{
					ID: "default-sidebar",
					Class: d.NewClass(
						"fixed",              // fixed position
						"top-0 left-0",       // from top left corner
						"z-40 w-64 h-screen", // position on top of screen witdth 64px and full screen height
						"transition-transform -translate-x-full sm:translate-x-0", // full screen for small screens
					),
					AriaLabel: "Sidebar",
				},
				d.Div(d.Props{
					Class: d.NewClass("h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800"),
				},
					d.UL(d.Props{Class: d.NewClass("space-y-2 font-medium")},
						d.LI(d.Props{},
							d.A(
								d.Props{Class: d.NewClass("flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group")},
								d.SVG(
									d.Props{
										Class:      d.NewClass("w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"),
										AriaHidden: "true",
										XMLNS:      "http://www.w3.org/2000/svg",
										Fill:       "currentColor",
										ViewBox:    "0 0 22 21",
									},
									d.Path(d.Props{D: "M16.975 11H10V4.025a1 1 0 0 0-1.066-.998 8.5 8.5 0 1 0 9.039 9.039.999.999 0 0 0-1-1.066h.002Z"}),
									d.Path(d.Props{D: "M12.5 0c-.157 0-.311.01-.565.027A1 1 0 0 0 11 1.02V10h8.975a1 1 0 0 0 1-.935c.013-.188.028-.374.028-.565A8.51 8.51 0 0 0 12.5 0Z"}),
								),
								d.Span(d.Props{InnerHTML: "Dashboard", Class: d.NewClass("ml-3")}),
							),
						),
					),
				),
			),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive
}
