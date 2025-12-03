"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
BoundsToRange ← □+⊙(⇡+₁)⟜-
Parse ← (
  ⊜□⊸≠@,
  ≡(⊜⋕⊸≠@-°□)
  ≡(BoundsToRange °⊟)
)
GenPowRange  ← ˜ⁿ 10 +₁⇡
GenDiffRange ← ⊟ ⊙+₁ GenPowRange⟜GenPowRange⌈÷2⊢△$"_"
FilterEven   ← ¬◿2⊢△$"_"
DivBar ← (
  ×⊸FilterEven
  ⌊°⊟÷⊸GenDiffRange
  ×⊃(≠0)=
  =1/+
)
Part₁ ← (
  ×⊸⍚(≡DivBar°□)
  °□/+⍚/+
)
Part₂ ← ∘
Main ← (
  ⊃Part₁Part₂ Parse
)
Main