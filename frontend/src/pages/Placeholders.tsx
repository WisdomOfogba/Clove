import React from 'react';
import { Link } from 'react-router-dom';

export function AboutPage() {
  return (
    <div className="min-h-[85vh] bg-white flex flex-col items-center justify-center p-8 text-center space-y-6">
      <h1 className="text-5xl font-black text-dark tracking-tighter">About CloveDelight</h1>
      <p className="max-w-2xl text-neutral font-medium text-lg leading-relaxed">
        CloveDelight is bringing radical trust and transparency to food delivery in Nigeria. Our platform leverages real-time AI and continuous video audits to guarantee that every meal prepared in a Clove-certified kitchen is handled with extreme hygiene standards.
      </p>
      <Link to="/" className="bg-dark text-white px-8 py-4 rounded-3xl font-black text-sm hover:bg-primary transition-colors">Return Home</Link>
    </div>
  );
}

export function HowItWorksPage() {
  return (
    <div className="min-h-[85vh] bg-white flex flex-col items-center justify-center p-8 text-center space-y-6">
      <h1 className="text-5xl font-black text-dark tracking-tighter">How Clove Works</h1>
      <div className="max-w-3xl space-y-4 text-left font-medium text-neutral mt-8">
         <div className="bg-slate-50 p-6 rounded-[2rem]"><strong className="text-dark hover:text-primary">1. AI Audited Kitchens:</strong> Every kitchen on Clove is constantly audited by our proprietary AI for cleanliness and standards.</div>
         <div className="bg-slate-50 p-6 rounded-[2rem]"><strong className="text-dark hover:text-primary">2. Secure Order & Prep:</strong> When you order, the kitchen prepares your meal under live surveillance which is logged for accountability.</div>
         <div className="bg-slate-50 p-6 rounded-[2rem]"><strong className="text-dark hover:text-primary">3. Geofenced Delivery:</strong> Our riders deliver the food securely, and the bag unlocks only when it arrives at your exact location.</div>
      </div>
      <Link to="/restaurants" className="mt-8 bg-dark text-white px-8 py-4 rounded-3xl font-black text-sm hover:bg-primary hover:text-dark transition-all active:scale-95 shadow-xl">Experience It</Link>
    </div>
  );
}

export function NotFoundPage() {
  return (
    <div className="min-h-[85vh] bg-slate-50 flex items-center justify-center p-8 text-center">
      <div className="space-y-6">
        <h1 className="text-9xl font-black text-dark tracking-tighter">404</h1>
        <p className="text-xl font-bold text-neutral">The kitchen is closed here.</p>
        <Link to="/" className="inline-block bg-primary text-dark px-8 py-4 rounded-3xl font-black text-sm hover:bg-accent transition-colors shadow-xl">
           Take Me Home
        </Link>
      </div>
    </div>
  );
}
