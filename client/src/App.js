
import "bootstrap/dist/css/bootstrap.min.css";
import Container from "react-bootstrap/Container";

function App() {
	return (
		<Container>
			<h1 className="text-center">react-mon-go</h1>
			<p>Baseline for a React, MongoDB, and Golang stack web application.</p>
			<h2>Setup</h2>
			<ol>
				<li>Fill in the environment variables in the <code>example.env</code></li>
				<li>Rename <code>example.env</code> to <code>.env</code></li>
			</ol>
			<h2>What&#39;s Included?</h2>
			<ul>
				<li>Golang<ul>
					<li>Gin</li>
				</ul>
				</li>
				<li>React<ul>
					<li>Bootstrap</li>
					<li>React-Bootstrap</li>
				</ul>
				</li>
				<li>MongoDB</li>
			</ul>
		</Container>
	);
}

export default App;
