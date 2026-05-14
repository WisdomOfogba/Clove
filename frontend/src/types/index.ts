export type UserRole = 'customer' | 'vendor' | 'admin';

export interface User {
  id: string;
  name: string;
  email: string;
  role: UserRole;
}

export interface BusinessInfo {
  name: string;
  address: string;
  phone: string;
  category: string;
}

export interface LegalDocs {
  cacNumber: string;
  ninNumber: string;
  cacFile: File | null;
  ninFile: File | null;
}

export interface ProofOfLife {
  kitchenPhotos: string[];
  mealPhotos: string[];
  exteriorPhoto: string | null;
  location: { lat: number; lng: number } | null;
  selfie: string | null;
}

export interface OnboardingData {
  step: number;
  business: BusinessInfo;
  legal: LegalDocs;
  proof: ProofOfLife;
  status: 'idle' | 'processing' | 'completed' | 'denied';
}

export interface TrustScore {
  score: number;
  verdict: 'verified' | 'restricted' | 'rejected';
  breakdown: {
    label: string;
    score: number;
    description: string;
  }[];
}
