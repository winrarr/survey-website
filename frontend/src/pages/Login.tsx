import type { Component } from 'solid-js';
import styles from './Signup.module.css';

const Home: Component = () => {
    return <>
        <iframe name="dummyframe" id="dummyframe" style="display: none;"></iframe>

        <div class={styles.container}>
            <form action="http://localhost:8080/login" method="post" class={styles.form} target="dummyframe">
                <div>
                    <label for="username">Username:</label>
                    <input type="text" name="username"></input><br></br>
                    <label for="password">Password:</label>
                    <input type="text" name="password"></input><br></br>
                </div>
                <input type="submit" value="Submit"></input>
            </form>
        </div>
    </>
};

export default Home;
