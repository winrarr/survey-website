import type { Component } from 'solid-js';
import {lazy} from 'solid-js';
import {Routes, Route, A } from '@solidjs/router';
import styles from './App.module.css';

const Home = lazy(() => import('./pages/Home'))
const Users = lazy(() => import('./pages/Users'))
const Signup = lazy(() => import('./pages/Signup'))
const Login = lazy(() => import('./pages/Login'))

const App: Component = () => {
  return <>
  <nav class={styles.navigation}>
    <A href='/'>Home</A>
    <A href='/login'>Login</A>
    <A href='/signup'>Sign Up</A>
    <A href='/users'>Users</A>
  </nav>
    <Routes>
      <Route path='/' component={Home}/>
      <Route path='/users' component={Users}/>
      <Route path='/signup' component={Signup}></Route>
      <Route path='/login' component={Login}></Route>
    </Routes>
  </>
};

export default App;
