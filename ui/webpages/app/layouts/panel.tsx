
import { Welcome } from "~/welcome/welcome";

export default function Panel() {
	return (
		<div className={"bg-gray-100 dark:bg-gray-800 h-dvh w-full grid grid-cols-24"} style={styles.Layout}>
			<div className="col-span-1">
				<div className="bg-transparent h-full p-32 flex items-center justify-between flex-left grid grid-rows-3">
						{/* Sidebar content can go here */}
						<span className="font-normal row-span-1">Panel</span>
				</div>
			</div>
			<div className="col-span-23">
				<nav className="bg-transparent min-h-[32px] top-0 left-0 w-full text-white justify-between items-center pt-3 pr-3 pb-2" >
					<div className="flex items-center justify-between flex-left grid grid-cols-3">
						<div className="flex items-center col-span-1 text-sm font-sans text-black dark:text-white gap-4">
							<span className="font-semibold">Snowbank</span>
							<span className="font-normal">Panel</span>
						</div>
						<div className="flex items-center col-span-1 justify-center text-sm font-sans text-black dark:text-white">
							{/* <span className="font-normal">Panel</span> */}
						</div>
						<div className="flex items-center col-span-1 justify-end text-sm font-sans text-black dark:text-white">
							{/* <span className="font-normal">Panel</span> */}
						</div>
					</div>
				</nav>
				<div className="bg-white w-full h-full rounded shadow-[0_4px_12px_0_rgba(0,0,0,0.1)]">
					<Welcome />
				</div>
			</div>

		</div>
	)
}

//add in file styles for the panel layout, do not import them, just add the styles
export const styles = {
	Layout: {
		// background: "rgb(239, 240, 243)",
		// background: 'rgb(25, 24, 31)',
		// background: 'rgb(215, 217, 224)'
	},
	Header: {
	},

};

