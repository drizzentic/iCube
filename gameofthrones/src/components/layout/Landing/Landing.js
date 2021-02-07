import React from 'react'
import { Link } from 'react-router-dom';

function Landing() {
    return (
        <div className="links">
            <Link to="/characters">Characters</Link>
            <Link to="/books">Books</Link>
            <Link to="/houses">Houses</Link>
        </div>
    )
}

export default Landing;
