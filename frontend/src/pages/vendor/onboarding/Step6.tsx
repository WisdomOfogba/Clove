import { Icon } from '@iconify/react';
import { motion, AnimatePresence } from 'motion/react';
import { useState } from 'react';
import { cn } from '../../../lib/utils';
import { useOnboardingStore } from '../../../store';

export function Step6() {
  const { reset } = useOnboardingStore();
  const [isExpanded, setIsExpanded] = useState(false);

  const trustData = {
    score: 92,
    verdict: 'verified' as const,
    breakdown: [
      { label: 'Environment Hygiene', score: 95, desc: 'Clutter-free and sanitized prep area detected.' },
      { label: 'Authenticity Match', score: 88, desc: 'Meals match Nigerian culinary markers for Jollof.' },
      { label: 'Identity Integrity', score: 100, desc: 'Facial recognition matches provided ID documents.' },
      { label: 'Location Accuracy', score: 90, desc: 'Geolocation matches business address within 5m.' }
    ]
  };

  return (
    <div className="p-8 space-y-8">
      <div className="flex flex-col items-center text-center space-y-4">
        {/* TrustScoreCircle */}
        <div className="relative">
          <svg className="w-40 h-40">
            <circle
              cx="80"
              cy="80"
              r="74"
              fill="none"
              stroke="#f1f5f9"
              strokeWidth="12"
            />
            <motion.circle
              cx="80"
              cy="80"
              r="74"
              fill="none"
              stroke="#6366F1"
              strokeWidth="12"
              strokeDasharray={2 * Math.PI * 74}
              initial={{ strokeDashoffset: 2 * Math.PI * 74 }}
              animate={{ strokeDashoffset: 2 * Math.PI * 74 * (1 - trustData.score / 100) }}
              transition={{ duration: 2, ease: "easeOut" }}
              strokeLinecap="round"
            />
          </svg>
          <div className="absolute inset-0 flex flex-col items-center justify-center">
            <motion.span 
              initial={{ opacity: 0, scale: 0.5 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ delay: 1 }}
              className="text-4xl font-black text-dark"
            >
              {trustData.score}%
            </motion.span>
            <span className="text-[10px] font-bold text-neutral uppercase tracking-widest leading-none">Trust Score</span>
          </div>
          <motion.div 
            initial={{ scale: 0 }}
            animate={{ scale: 1 }}
            transition={{ delay: 1.5, type: "spring" }}
            className="absolute -bottom-2 -right-2 bg-primary text-white p-2 rounded-full shadow-lg shadow-primary/30"
          >
            <Icon icon="lucide:shield-check" className="w-6 h-6" />
          </motion.div>
        </div>

        <div className="space-y-1">
          <motion.div
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 1.2 }}
            className="inline-flex items-center gap-2 bg-primary/10 text-primary px-4 py-1.5 rounded-full text-sm font-bold border border-primary/20 mb-2"
          >
            <Icon icon="lucide:award" className="w-4 h-4" /> Verdict: Certified Premium Vendor
          </motion.div>
          <h2 className="text-2xl font-bold tracking-tight text-dark">Welcome to CloveDelight!</h2>
          <p className="text-sm text-neutral max-w-sm">
            Your verification was successful. Your kitchen is now globally visible to thousands of customers.
          </p>
        </div>
      </div>

      {/* AI Breakdown Accordion */}
      <div className="bg-slate-50 rounded-2xl border border-slate-100 overflow-hidden">
        <button 
          onClick={() => setIsExpanded(!isExpanded)}
          className="w-full flex items-center justify-between p-4 hover:bg-slate-100 transition-colors"
        >
          <div className="flex items-center gap-2">
            <Icon icon="lucide:alert-circle" className="w-4 h-4 text-primary" />
            <span className="text-xs font-bold text-neutral uppercase tracking-wider">Detailed AI Breakdown</span>
          </div>
          <Icon icon="lucide:chevron-down" className={cn("w-4 h-4 text-neutral transition-transform", isExpanded && "rotate-180")} />
        </button>
        
        <AnimatePresence>
          {isExpanded && (
            <motion.div
              initial={{ height: 0 }}
              animate={{ height: "auto" }}
              exit={{ height: 0 }}
              className="border-t border-slate-100"
            >
              <div className="p-4 space-y-4">
                {trustData.breakdown.map((item, i) => (
                  <div key={i} className="space-y-2">
                    <div className="flex items-center justify-between text-xs font-bold">
                      <span className="text-neutral">{item.label}</span>
                      <span className="text-primary">{item.score}%</span>
                    </div>
                    <div className="h-1.5 bg-slate-200 rounded-full overflow-hidden">
                      <motion.div 
                        initial={{ width: 0 }}
                        animate={{ width: `${item.score}%` }}
                        className="h-full bg-primary"
                      />
                    </div>
                    <p className="text-[10px] text-neutral/70 leading-relaxed">{item.desc}</p>
                  </div>
                ))}
              </div>
            </motion.div>
          )}
        </AnimatePresence>
      </div>

      <div className="space-y-3">
        <button 
          className="w-full bg-primary text-white rounded-xl py-4 font-bold hover:opacity-90 transition-all flex items-center justify-center gap-2 shadow-lg shadow-primary/20 group"
        >
          Setup Your Digital Menu
          <Icon icon="lucide:arrow-right" className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
        </button>
        <div className="grid grid-cols-2 gap-3">
          <button className="bg-white text-neutral border border-slate-200 rounded-xl py-3 text-sm font-bold hover:bg-slate-50 transition-all flex items-center justify-center gap-2">
             <Icon icon="lucide:share-2" className="w-4 h-4" /> Share Status
          </button>
          <button 
            onClick={reset}
            className="bg-white text-neutral border border-slate-200 rounded-xl py-3 text-sm font-bold hover:bg-slate-50 transition-all"
          >
             Onboarding Demo Reset
          </button>
        </div>
      </div>
    </div>
  );
}
