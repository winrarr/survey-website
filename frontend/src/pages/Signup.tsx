import type { Component, createResource } from 'solid-js';
import { createStore } from 'solid-js/store'
import styles from './Signup.module.css';

type FormFields = {
    username?: string;
    password?: string;
    name?: string;
}

const Home: Component = () => {
    const [form, setForm] = createStore<FormFields>();

    const handleSubmit = async (event: Event): Promise<void> => {
        event.preventDefault();
        
        const dataToSubmit = {
            username: form.username,
            password: form.password,
            name: form.name,
        };

        console.log(JSON.stringify(dataToSubmit));
        
        let session = (await fetch('http://localhost:8080/user', {
            method: 'POST',
            mode: 'no-cors',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(dataToSubmit)
        })).text
        
        console.log(session)
    }

    const updateFormField = (fieldName: string) => (event: Event) => {
        const inputElement = event.currentTarget as HTMLInputElement;
        setForm({[fieldName]: inputElement.value})
    }

    return <>
        <iframe name="dummyframe" id="dummyframe" style="display: none;"></iframe>

        <div class={styles.container}>
            {/* <form action="http://localhost:8080/user" method="post" class={styles.form} target="dummyframe">
                <div>
                    <label for="username">Username:</label>
                    <input type="text" name="username"></input><br></br>
                    <label for="password">Password:</label>
                    <input type="text" name="password"></input><br></br>
                    <label for="name">Name:</label>
                    <input type="text" name="name"></input><br></br>
                </div>
                <input type="submit" value="Submit"></input>
            </form> */}
            <form class={styles.form} onsubmit={handleSubmit}>
                <div>
                    <label for="username">Username:</label>
                    <input onchange={updateFormField("username")} type="text"></input><br></br>
                    <label for="password">Password:</label>
                    <input onchange={updateFormField("password")} type="text"></input><br></br>
                    <label for="name">Name:</label>
                    <input onchange={updateFormField("name")} type="text"></input><br></br>
                </div>
                <input type="submit" value="Submit"></input>
            </form>
        </div>
    </>

    // return <>
    //     <div class={styles.container}>
    //         <div class={styles.loginContainer}>
    //             <h1 class={styles.loginText}>Login.</h1>
    //                 <input
    //                     id="email"
    //                     class="mt-2 border p-2 rounded-sm"
    //                     type="email"
    //                     placeholder="Your email"
    //                 ></input>
    //             <button>Send Magic Link</button>
    //         </div>
    //     </div>
    // </>
};

export default Home;
