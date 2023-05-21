import auth from '@react-native-firebase/auth';
import {
  GoogleSignin,
  User,
  statusCodes
} from '@react-native-google-signin/google-signin';
import { createContext, useContext, useEffect } from 'react';
import React, { useState } from 'react';

import { WEB_CLIENT_ID } from '../../../../environment';

GoogleSignin.configure({
  webClientId: WEB_CLIENT_ID
});

const AuthContext = createContext({} as User | null);

export const useAuth = () => {
  const [userInfo, setUserInfo] = useState<User | null>(null);
  const [isSigninInProgress, setIsSigninInProgress] = useState<boolean>(false);

  useEffect(() => {
    const subscriber = auth().onAuthStateChanged(async (_user) => {
      if (_user) {
        const currentUser = await GoogleSignin.getCurrentUser();
        setUserInfo(currentUser);
      } else {
        setUserInfo(null);
      }
    });

    return subscriber;
  }, []);

  //Googleログイン
  const onGoogleLogin = async () => {
    setIsSigninInProgress(true);
    try {
      await GoogleSignin.hasPlayServices({
        showPlayServicesUpdateDialog: true
      });
      const userInfo = await GoogleSignin.signIn();
      setUserInfo(userInfo);
    } catch (error: any) {
      if (error.code === statusCodes.SIGN_IN_CANCELLED) {
        // console.log('SIGN_IN_CANCELLED');
      } else if (error.code === statusCodes.IN_PROGRESS) {
        // console.log('IN_PROGRESS');
      } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
        // console.log('PLAY_SERVICES_NOT_AVAILABLE');
      } else {
        console.log(error);
      }
    }
    setIsSigninInProgress(false);
  };

  //サインアウト
  const onSignOut = async () => {
    try {
      await GoogleSignin.signOut();
      setUserInfo(null);
    } catch (error: any) {
      console.error(error);
    }
  };

  return { userInfo, isSigninInProgress, onGoogleLogin, onSignOut };
};

export const useAuthContext = () => {
  return useContext(AuthContext);
};

type Props = {
  children: React.ReactNode;
};

export const AuthProvider = ({ children }: Props) => {
  const { userInfo } = useAuth();

  return (
    <AuthContext.Provider value={userInfo}>{children}</AuthContext.Provider>
  );
};
