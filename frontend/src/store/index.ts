import { create } from 'zustand';
import { OnboardingData, UserRole } from '../types';

interface AuthState {
  user: { name: string; email: string; role: UserRole } | null;
  setUser: (user: { name: string; email: string; role: UserRole } | null) => void;
  setRole: (role: UserRole) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: { name: 'Demo User', email: 'user@example.com', role: 'customer' },
  setUser: (user) => set({ user }),
  setRole: (role) => set((state) => ({ 
    user: state.user ? { ...state.user, role } : null 
  })),
}));

interface OnboardingState {
  data: OnboardingData;
  setStep: (step: number) => void;
  updateBusiness: (data: Partial<OnboardingState['data']['business']>) => void;
  updateLegal: (data: Partial<OnboardingState['data']['legal']>) => void;
  updateProof: (data: Partial<OnboardingState['data']['proof']>) => void;
  setStatus: (status: OnboardingData['status']) => void;
  reset: () => void;
}

const initialOnboardingData: OnboardingData = {
  step: 1,
  business: { name: '', address: '', phone: '', category: '' },
  legal: { cacNumber: '', ninNumber: '', cacFile: null, ninFile: null },
  proof: { brandImages: [], brandLogo: null, location: null, selfie: null, identityReference: null },
  status: 'idle',
};

export const useOnboardingStore = create<OnboardingState>((set) => ({
  data: initialOnboardingData,
  setStep: (step) => set((state) => ({ data: { ...state.data, step } })),
  updateBusiness: (business) => set((state) => ({ data: { ...state.data, business: { ...state.data.business, ...business } } })),
  updateLegal: (legal) => set((state) => ({ data: { ...state.data, legal: { ...state.data.legal, ...legal } } })),
  updateProof: (proof) => set((state) => ({ data: { ...state.data, proof: { ...state.data.proof, ...proof } } })),
  setStatus: (status) => set((state) => ({ data: { ...state.data, status } })),
  reset: () => set({ data: initialOnboardingData }),
}));
