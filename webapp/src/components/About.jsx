import React from 'react';
import {Link} from "@reach/router";

function About() {
	return (<>
		<h2><Link to="/about">About Gas Giant Battles</Link></h2>
		<p className="introtext">
			Gas Giant Battles is a strategic, turn-based game of conquest and accounting.
		</p>
		<p>
			Players begin the game in control of one star-system.
			They mine stuff.
			They build stuff.
			They make ships.
			They head to the stars.
		</p>

		<h2><a href="https://github.com/mdhender/server/LICENSE">License</a></h2>
		<p>
			Gas Giant Battles is licensed under the GNU Affero General Public License, version 3.
		</p>
		<code>
			<p>
				This program is free software: you can redistribute it and/or modify
				it under the terms of the GNU Affero General Public License as published
				by the Free Software Foundation, either version 3 of the License, or
				(at your option) any later version.
			</p>
			<p>
				This program is distributed in the hope that it will be useful,
				but WITHOUT ANY WARRANTY; without even the implied warranty of
				MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
				GNU Affero General Public License for more details.
			</p>
			<p>
				You should have received a copy of the GNU Affero General Public License
				along with this program. If not, see <a href="https://www.gnu.org/licenses/">licenses</a>.
			</p>
		</code>
	</>);
}

export default About;
