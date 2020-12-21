import React from 'react';
import {useSelector} from 'react-redux';
import {selectCopyright} from '../features/site/siteSlice';
import {Link} from "@reach/router";

function Footer() {
	const copyright = useSelector(selectCopyright);

	return (
			<div id="footer">
				<div className="left">
					<p>&copy; {copyright.year} {copyright.author} | <a href="https://github.com/mdhender/server">Gas Giant Battles</a> | Template design by <a href="https://github.com/mdhender/server/VIKLUND.md">Andreas Viklund</a></p>
				</div>
				<div className="right textright">
					<p><Link to="/about">About</Link> | <a href="https://github.com/mdhender/server">Source</a></p>
					<p className="hide"><a href="#top">Return to top</a></p>
				</div>
			</div>
	);
}

export default Footer;