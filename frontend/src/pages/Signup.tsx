import type { Component } from 'solid-js';

const Home: Component = () => {
    return <>
        <div class="container">
            <div class="login-container">
                <h1 class="login-text">Login.</h1>
                    <input
                        id="email"
                        class="mt-2 border p-2 rounded-sm"
                        type="email"
                        placeholder="Your email"
                    ></input>
                <button>Send Magic Link</button>
            </div>
        </div>
    </>
};

export default Home;
