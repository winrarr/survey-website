import type { Component } from 'solid-js';

const Home: Component = () => {
    return <>
        <div class="flex mt-20 items-center justify-center">
            <div class="flex flex-col border p-4 rounded-lg shadow-lg w-64">
                <h1 class="text-2xl">Login.</h1>
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
