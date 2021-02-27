<script>
	import logo from "./logo.png";
	import {Tabs} from "attractions"
	import { navigate , Router, Link, Route } from "svelte-routing";
	import DistressedInfo from "./components/DistressedInfo.svelte";
	import WeatherInfo from "./components/WeatherInfo.svelte";
	import SARInfo from "./components/SARInfo.svelte";
	import Result from "./components/Result.svelte";

	let selectedTab;
	const handleTabs= (changedValue) => {
		let navTo ="";
		switch (changedValue.detail.value) {
			case 'Distressed Unit Information':
			navTo=	'DistressedInfo';
				break;
				case 'Weather Information':
			navTo=	'WeatherInfo';
				break;
				case 'SAR Unit Information':
			navTo=	'SARInfo';
				break;
				case 'Result':
			navTo=	'Result';
				break;
		
	
		}
		console.log(`navto: ${navTo}		changedValue: ${changedValue.detail.value}` );
	navigate(`/${navTo}`);
	};
export let url='';
</script>

<main>
	<div class="App">
		<header class="App-header">
			<Router url="{url}">
			<Tabs on:change={handleTabs}

			
				name="Nav"
				items={['Distressed Unit Information', 
				'Weather Information', 'SAR Unit Information','Result']}
				bind:value={selectedTab}

			/>
			<div>
				<Route path="DistressedInfo" component="{DistressedInfo}" />
				<Route path="WeatherInfo" component="{WeatherInfo}" />
				<Route path="SARInfo" component="{SARInfo}" />
				<Route path="Result" component="{Result}" />
			
			  </div>
		</Router>

		</header>
	</div>
</main>

<style>
	:global(body) {
		margin: 0;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
			"Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans",
			"Helvetica Neue", sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	code {
		font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
			monospace;
	}

	.App {
		text-align: center;
	}

	

	.App-header {
		background-color: #282c34;
		/*min-height: 100vh;*/
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		font-size: calc(10px + 2vmin);
		color: white;
		flex-grow: 1;
		
	}


	@keyframes App-logo-spin {
		from {
			transform: rotate(0deg);
		}
		to {
			transform: rotate(360deg);
		}
	}
</style>
