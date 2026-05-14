import { Icon } from '@iconify/react';
import { motion } from 'motion/react';
import { useState, useEffect } from 'react';

export function Step5() {
  const [status, setStatus] = useState('Initializing Clove AI...');
  const [progress, setProgress] = useState(0);

  const messages = [
    'Scanning kitchen environmental photos...',
    'Analyzing meal authenticity markers...',
    'Performing biometric face match...',
    'Verifying geolocation coordinates...',
    'Validating CAC legal database records...',
    'Calculating Trust Score...',
  ];

  useEffect(() => {
    let currentIdx = 0;
    const interval = setInterval(() => {
      if (currentIdx < messages.length) {
        setStatus(messages[currentIdx]);
        currentIdx++;
        setProgress(prev => prev + 15);
      } else {
        clearInterval(interval);
      }
    }, 700);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="p-12 flex flex-col items-center justify-center text-center space-y-8 min-h-[500px]">
      <div className="relative">
        <motion.div
           animate={{ rotate: 360 }}
           transition={{ duration: 4, repeat: Infinity, ease: "linear" }}
           className="w-32 h-32 rounded-full border-4 border-slate-100 border-t-primary"
        />
        <div className="absolute inset-0 flex items-center justify-center">
          <Icon icon="lucide:shield-check" className="w-12 h-12 text-primary animate-pulse" />
        </div>
      </div>

      <div className="space-y-4 max-w-xs">
        <h2 className="text-2xl font-bold tracking-tight text-dark">Verifying Life Proof</h2>
        <div className="h-1.5 w-full bg-slate-100 rounded-full overflow-hidden">
          <motion.div 
            initial={{ width: 0 }}
            animate={{ width: `${progress}%` }}
            className="h-full bg-primary"
          />
        </div>
        <p className="text-sm font-medium text-neutral animate-pulse">{status}</p>
      </div>

      <div className="pt-8 border-t border-slate-50 w-full grid grid-cols-2 gap-4">
        <div className="flex items-center gap-3 bg-slate-50 p-3 rounded-xl">
           <Icon icon="lucide:shield-check" className="w-5 h-5 text-primary" />
           <div className="text-left">
             <p className="text-[10px] font-bold text-neutral uppercase tracking-wider">Security</p>
             <p className="text-xs font-bold text-dark">Encrypted</p>
           </div>
        </div>
        <div className="flex items-center gap-3 bg-slate-50 p-3 rounded-xl">
           <Icon icon="lucide:loader-2" className="w-5 h-5 text-accent animate-spin" />
           <div className="text-left">
             <p className="text-[10px] font-bold text-neutral uppercase tracking-wider">Analysis</p>
             <p className="text-xs font-bold text-dark">Deep Neural</p>
           </div>
        </div>
      </div>
    </div>
  );
}
