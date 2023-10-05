grammar Grammar;


s returns [[]interface{} code]
: block EOF
    {
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction*
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;


arguments returns [[]interface{} args]
@init {
    $args = []interface{}{}
}
    :argument COMA arguments  { $args = append($args, $argument.e)
                                       for _, arg := range $arguments.args {
                                           $args = append($args, arg)
                                       }
                                 }
    |argument { $args = append( $args , $argument.e) }
    | {}
    ;

argument returns [interface{} e]
    : expr { $e = $expr.e; }
    ;

guardstatement returns [interfaces.Instruction newguard]
: RGUARD expr RELSE LLAVEIZQ block LLAVEDER {$newguard = instructions.NewGuard($RGUARD.line, $RGUARD.pos,$expr.e, $block.blk )}
;



// INSTRUCTIONS
instruction returns [interfaces.Instruction inst]
: printstmt PTOCOMA?  { $inst = $printstmt.prnt}
| structfuncall   {$inst = $structfuncall.newstructfunccall}
| vecdec PTOCOMA? {$inst = $vecdec.newvecdec}
| vardec PTOCOMA?  { $inst = $vardec.newdec}
| constdec PTOCOMA? {$inst = $constdec.newconst}
| appendvec PTOCOMA? {$inst = $appendvec.newappendvec}
| decmatrix PTOCOMA? {$inst = $decmatrix.newmatrix}
| removelastvec PTOCOMA? {$inst = $removelastvec.newremovelastvec}
| removeatvec PTOCOMA?  {$inst = $removeatvec.newremoveat}
| asignation PTOCOMA? {$inst = $asignation.newasignation}
| unarysum PTOCOMA?  {$inst = $unarysum.newunarysum}
| unarysub PTOCOMA?  {$inst = $unarysub.newunarysub}
| breakstatement PTOCOMA? {$inst = $breakstatement.newbreak}
| continuestatement PTOCOMA? {$inst = $continuestatement.newcontinue}
| vectormodification PTOCOMA? {$inst = $vectormodification.newvecmod}
| ifstmt {$inst = $ifstmt.newif}
| while_statement {$inst = $while_statement.newwhile}
| switchstatement {$inst = $switchstatement.newswitch}
| forloop {$inst = $forloop.newfor}
| structdef {$inst = $structdef.newstruct}
| funcdec   {$inst = $funcdec.newfuncdec}
| retturn PTOCOMA? {$inst = $retturn.newreturn}
| callfuncins PTOCOMA? {$inst = $callfuncins.newcallfuncins}
| structmodification PTOCOMA? {$inst = $structmodification.newstructmod}
| guardstatement  {$inst = $guardstatement.newguard}
;


// INSTRUCTIONS

structfuncall returns [interfaces.Instruction newstructfunccall]
: id1=ID PTO ID  PARIZQ PARDER  {$newstructfunccall = instructions.NewStructFuncCall($id1.line, $id1.pos, $id1.text, $ID.text)}
;


structinstruction returns [interfaces.Instruction inst]
: vecdec PTOCOMA? {$inst = $vecdec.newvecdec}
| vardec PTOCOMA?  { $inst = $vardec.newdec}
| constdec PTOCOMA? {$inst = $constdec.newconst}
| decmatrix PTOCOMA? {$inst = $decmatrix.newmatrix}
| structfuncdec {$inst = $structfuncdec.newfuncdec}
;

structfuncdec returns [interfaces.Instruction newfuncdec]
: RFUNC ID PARIZQ funcparameterlist  PARDER ARROW typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) LLAVEIZQ funcblock LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, $typpe.text,$funcblock.blk, false )}
| RFUNC ID PARIZQ funcparameterlist  PARDER  LLAVEIZQ funcblock LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, nil,$funcblock.blk, false )}
| RMUTATING RFUNC ID PARIZQ funcparameterlist  PARDER ARROW typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) LLAVEIZQ funcblock LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, $typpe.text,$funcblock.blk, true )}
| RMUTATING RFUNC ID PARIZQ funcparameterlist  PARDER  LLAVEIZQ funcblock LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, nil,$funcblock.blk, true )}
;




funcblock returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IFuncinstContext
  }
: ins+=funcinst*
    {
        listInt = localctx.(*FuncblockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

 funcinst returns [interfaces.Instruction inst]
: printstmt PTOCOMA?  { $inst = $printstmt.prnt}
   | vecdec PTOCOMA? {$inst = $vecdec.newvecdec}
   | vardec PTOCOMA?  { $inst = $vardec.newdec}
   | constdec PTOCOMA? {$inst = $constdec.newconst}
   | appendvec PTOCOMA? {$inst = $appendvec.newappendvec}
   | decmatrix PTOCOMA? {$inst = $decmatrix.newmatrix}
   | removelastvec PTOCOMA? {$inst = $removelastvec.newremovelastvec}
   | removeatvec PTOCOMA?  {$inst = $removeatvec.newremoveat}
   | asignation PTOCOMA? {$inst = $asignation.newasignation}
   | unarysum PTOCOMA?  {$inst = $unarysum.newunarysum}
   | unarysub PTOCOMA?  {$inst = $unarysub.newunarysub}
   | breakstatement PTOCOMA? {$inst = $breakstatement.newbreak}
   | continuestatement PTOCOMA? {$inst = $continuestatement.newcontinue}
   | vectormodification PTOCOMA? {$inst = $vectormodification.newvecmod}
   | ifstmt {$inst = $ifstmt.newif}
   | while_statement {$inst = $while_statement.newwhile}
   | switchstatement {$inst = $switchstatement.newswitch}
   | forloop {$inst = $forloop.newfor}
   | callfuncins PTOCOMA? {$inst = $callfuncins.newcallfuncins}
   | guardstatement  {$inst = $guardstatement.newguard}
   | selfstructmodification {$inst = $selfstructmodification.newstructmod}
   ;



selfstructmodification returns [interfaces.Instruction newstructmod]
: RSELF PTO ID IG expr  {$newstructmod = instructions.NewSelfModification($RSELF.line, $RSELF.pos, $ID.text, $expr.e)}

;





structmodification returns [interfaces.Instruction newstructmod]
: ID PTO attrlist IG expr  {$newstructmod = instructions.NewStructMod($ID.line, $ID.pos, $ID.text, $attrlist.atrlist, $expr.e)}

;


structblock returns [[]interface{} sblk]
@init{
    $sblk = []interface{}{}
    var listsinst []IStructinstructionContext
  }
: ins+=structinstruction*
    {
        listsinst = localctx.(*StructblockContext).GetIns()
        for _, e := range listsinst {
            $sblk = append($sblk, e.GetInst())
        }
    }
;

// TODO ADD FUNCTIONS TO STRUCTS

structdef returns [interfaces.Instruction newstruct]
: RSTRUCT ID LLAVEIZQ structblock LLAVEDER {$newstruct = instructions.NewStructDef($ID.line, $ID.pos,$ID.text ,$structblock.sblk)}
;



retturn returns [interfaces.Instruction newreturn] :

RRETURN expr {$newreturn = instructions.NewReturn($RRETURN.line, $RRETURN.pos,$expr.e )}
| RRETURN  {$newreturn = instructions.NewReturn($RRETURN.line, $RRETURN.pos,nil)}

;


funcdec returns [interfaces.Instruction newfuncdec]
: RFUNC ID PARIZQ funcparameterlist  PARDER ARROW typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) LLAVEIZQ block LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, $typpe.text,$block.blk, false )}
| RFUNC ID PARIZQ funcparameterlist  PARDER  LLAVEIZQ block LLAVEDER {$newfuncdec = instructions.NewFuncDec($RFUNC.line, $RFUNC.pos,$ID.text ,$funcparameterlist.fplist, nil,$block.blk, false )}

;


vectormodification returns [interfaces.Instruction newvecmod]
: ID indexesList IG expr {$newvecmod = instructions.NewVectorModification($ID.line, $ID.pos,$ID.text ,$indexesList.indexes, $expr.e)}

;

forloop returns [interfaces.Instruction newfor]
: RFOR ID RIN ex=expr LLAVEIZQ b=block LLAVEDER {$newfor = instructions.NewFor($RFOR.line, $RFOR.pos,$ID.text ,$ex.e, $b.blk)}
| RFOR ID RIN range LLAVEIZQ b=block LLAVEDER {$newfor = instructions.NewFor($RFOR.line, $RFOR.pos,$ID.text ,$range.newrange, $b.blk)}

;

range returns [interfaces.Expression newrange]
: exp1=expr PTO PTO PTO exp2=expr {$newrange = expressions.NewRange($exp1.start.GetLine(), $exp1.start.GetColumn(), $exp1.e, $exp2.e) }

;

//--------------------------
removeatvec returns [interfaces.Instruction newremoveat]
: ID PTO RREMOVEAT PARIZQ RRAT DOSPTOS expr PARDER {$newremoveat = instructions.NewRemoveAtVector($ID.line, $ID.pos,$ID.text ,$expr.e)}
  ;

appendvec returns [interfaces.Instruction newappendvec]
: ID PTO RAPPEND PARIZQ expr PARDER {$newappendvec = instructions.NewAppendVector($ID.line, $ID.pos,$ID.text ,$expr.e)}

;

removelastvec returns [interfaces.Instruction newremovelastvec]
: ID PTO RREMOVELAST PARIZQ PARDER {$newremovelastvec = instructions.NewRemoveLastVector($ID.line, $ID.pos,$ID.text)}

;


vecdec returns [interfaces.Instruction newvecdec]
: RVAR ID DOSPTOS OBRA typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) CBRA IG OBRA typpe2=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) CBRA OBRA CBRA {$newvecdec = instructions.NewVecDec($RVAR.line, $RVAR.pos, $ID.text, $typpe.text, $typpe2.text, nil )}
| RVAR firstid=ID DOSPTOS  OBRA typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) CBRA IG expr {$newvecdec = instructions.NewVecDec($RVAR.line, $RVAR.pos, $firstid.text, $typpe.text, nil,$expr.e )}
| RVAR id1 = ID IG OBRA id2=ID CBRA PARIZQ PARDER {$newvecdec = instructions.NewVecDec($RVAR.line, $RVAR.pos, $id1.text, $id2.text, nil, nil )}
;

breakstatement returns [interfaces.Instruction newbreak]
: RBREAK    {$newbreak = instructions.NewBreak($RBREAK.line, $RBREAK.pos)}

;

continuestatement returns [interfaces.Instruction newcontinue]
: RCONTINUE    {$newcontinue = instructions.NewContinue($RCONTINUE.line, $RCONTINUE.pos)}

;

switchstatement returns [interfaces.Instruction newswitch]
: RSWITCH  ex=expr  LLAVEIZQ caselist LLAVEDER   {$newswitch = instructions.NewSwitch($RSWITCH.line, $RSWITCH.pos, $ex.e, $caselist.newcaselist, nil)}
 |RSWITCH  ex=expr   LLAVEIZQ caselist defaultstatement LLAVEDER   {$newswitch = instructions.NewSwitch($RSWITCH.line, $RSWITCH.pos, $ex.e, $caselist.newcaselist, $defaultstatement.newdefault)}
;


caselist returns [[]interface{} newcaselist]
@init {
    $newcaselist = []interface{}{}
}
    :case  caselist  { $newcaselist = append($newcaselist, $case.newcase)
                                       for _, arg := range $caselist.newcaselist {
                                           $newcaselist = append($newcaselist, arg)
                                       }
                                 }
    |case { $newcaselist = append( $newcaselist , $case.newcase) }
    | {}
    ;

case returns [interfaces.Instruction newcase]
: RCASE ex=expr DOSPTOS  b=block   {$newcase = instructions.NewCase($RCASE.line, $RCASE.pos, $ex.e, $b.blk)}
;

defaultstatement returns [[] interface{} newdefault]
: RDEFAULT DOSPTOS b=block  LLAVEDER             {$newdefault =  $b.blk}
;

//-----------------------------

ifstmt returns [interfaces.Instruction newif]
: RIF  ex=expr  LLAVEIZQ b=block LLAVEDER eliflist  {$newif = instructions.NewIf($RIF.line, $RIF.pos, $ex.e, $b.blk, $eliflist.neweliflist, nil)}
 |RIF  ex=expr  LLAVEIZQ b=block LLAVEDER eliflist elsestament {$newif = instructions.NewIf($RIF.line, $RIF.pos, $ex.e, $b.blk, $eliflist.neweliflist, $elsestament.newelse)}
;


eliflist returns [[]interface{} neweliflist]
@init {
    $neweliflist = []interface{}{}
}
    :elif  eliflist  { $neweliflist = append($neweliflist, $elif.newelif)
                                       for _, arg := range $eliflist.neweliflist {
                                           $neweliflist = append($neweliflist, arg)
                                       }
                                 }
    |elif { $neweliflist = append( $neweliflist , $elif.newelif) }
    | {}
    ;

elif returns [interfaces.Instruction newelif]
: RELSE RIF ex=expr LLAVEIZQ  b=block  LLAVEDER {$newelif = instructions.NewIf($RELSE.line, $RELSE.pos, $ex.e, $b.blk, nil, nil)}
;

elsestament returns [[] interface{} newelse]
: RELSE LLAVEIZQ b=block  LLAVEDER             {$newelse =  $b.blk}
;
// STATEMENTS----------------------------------------------------------------------------------------------
printstmt returns [interfaces.Instruction prnt]
: RPRINT PARIZQ arguments PARDER { $prnt = instructions.NewPrint($RPRINT.line,$RPRINT.pos,$arguments.args)}
;

while_statement returns [interfaces.Instruction newwhile]
: RWHILE expr LLAVEIZQ b=block LLAVEDER { $newwhile = instructions.NewWhile($RWHILE.line,$RWHILE.pos,$expr.e, $b.blk)}
;
vardec returns [interfaces.Instruction newdec]
: RVAR ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) IG ex=expr { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text,$typpe.text, $ex.e)}
| RVAR ID IG ex=expr                                                      { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text, nil, $ex.e)}
| RVAR ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) QM         { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text, $typpe.text, nil)}
;




constdec returns [interfaces.Instruction newconst]
: RLET ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) IG ex=expr { $newconst = instructions.NewConstDec($RLET.line,$RLET.pos,$ID.text,$typpe.text, $ex.e)}
| RLET ID IG ex=expr                                                      { $newconst = instructions.NewConstDec($RLET.line,$RLET.pos,$ID.text, nil, $ex.e)}
;

asignation returns [interfaces.Instruction newasignation]
: ID IG ex=expr { $newasignation = instructions.NewAsignation($ID.line,$ID.pos,$ID.text, $ex.e)}
;




unarysum returns [interfaces.Instruction newunarysum]
:ID UNARYPLUS ex=expr { $newunarysum = instructions.NewUnarySum($ID.line,$ID.pos,$ID.text, "+=", $ex.e)}
;

unarysub returns [interfaces.Instruction newunarysub]
:ID UNARYMINUS ex=expr { $newunarysub = instructions.NewUnarySum($ID.line,$ID.pos,$ID.text, "-=", $ex.e)}
;

isemptyvec returns [interfaces.Expression newisemptyvec]
: ID PTO RISEMPTY   {$newisemptyvec = expressions.NewIsEmptyVector($ID.line, $ID.pos,$ID.text)}
  ;

countvec returns [interfaces.Expression newcountvec]
: ID PTO RCOUNT   {$newcountvec = expressions.NewCountVector($ID.line, $ID.pos,$ID.text)}
   ;



//-----------------------
vectoraccess returns [interfaces.Expression newvecaccess]
: ID indexesList  {$newvecaccess = expressions.NewVectorAccess($ID.line, $ID.pos,$ID.text, $indexesList.indexes)}
;

structaccess returns [interfaces.Expression saccess]
: ID PTO attrlist  {$saccess = expressions.NewStructAccess($ID.line, $ID.pos, $ID.text, $attrlist.atrlist)}

;

vecindexstruct returns  [[]string vecinlist]
: PTO attrlist  {$vecinlist = $attrlist.atrlist}


;

vectorstructaccess returns [interfaces.Expression vecstructaccess]:
ID indexesList vecindexstruct {$vecstructaccess = expressions.NewVectorStructAccess($ID.line, $ID.pos,$ID.text, $indexesList.indexes, $vecindexstruct.vecinlist)}
;



indexesList returns [[]interface{} indexes]
@init {
    $indexes = []interface{}{}
}
: vecac  i=indexesList {            $indexes = append($indexes, $vecac.newvecac)
                                       for _, arg := range $i.indexes{
                                           $indexes = append($indexes, arg)
                                      }

                                 }
    |vecac { $indexes = append( $indexes , $vecac.newvecac) }
    ;

vecac returns [interface{} newvecac]
    : OBRA expr CBRA { $newvecac = $expr.e; }
;



matrix_type returns [string newmatrixtype]
: OBRA typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) CBRA {$newmatrixtype = $OBRA.text + $typpe.text + $CBRA.text}
| OBRA matrix_type CBRA   {$newmatrixtype = $OBRA.text + $matrix_type.text + $CBRA.text}
;

repeatingvector returns [interfaces.Expression newrepeatingvec]
: matrix_type PARIZQ RREPEATING DOSPTOS r=repeatingvector COMA RCOUNT DOSPTOS expr PARDER  {$newrepeatingvec = expressions.NewRepeatingVector($r.start.GetLine(), $r.start.GetColumn(),$matrix_type.text,$repeatingvector.newrepeatingvec,$expr.e)}
| matrix_type PARIZQ RREPEATING DOSPTOS ex1=expr COMA RCOUNT DOSPTOS ex2=expr PARDER  {$newrepeatingvec = expressions.NewRepeatingVector($matrix_type.start.GetLine(), $matrix_type.start.GetColumn(),$matrix_type.text, $ex1.e,$ex2.e)}

;

manualdef   returns [interfaces.Expression newmanualdef]
: manualmatrixdef {$newmanualdef = expressions.NewManualMatrixDef($manualmatrixdef.start.GetLine(), $manualmatrixdef.start.GetColumn(),$manualmatrixdef.newmanualmatrixdef ) }

;


manualmatrixdef returns [[]interface {} newmanualmatrixdef]

: OBRA values2 CBRA {$newmanualmatrixdef = append($newmanualmatrixdef, $values2.newvalueslist)
                    }
;




values2 returns [[]interface{} newvalueslist]
@init {
    $newvalueslist = []interface{}{}
}
: v=values2 COMA manualmatrixdef {
    $newvalueslist = append($v.newvalueslist, $manualmatrixdef.newmanualmatrixdef...)
}
| manualmatrixdef {
    $newvalueslist = append($newvalueslist, $manualmatrixdef.newmanualmatrixdef...)
}
| arguments {
    $newvalueslist = append($newvalueslist, $arguments.args...)
}
;

funcarglist returns [[]environment.FuncArg fplist]
@init {
    $fplist = make([]environment.FuncArg, 0)
}
     : funcarg { $fplist = append($fplist, $funcarg.fp) }
       COMA a=funcarglist  { $fplist = append($fplist, $a.fplist...) }
       | funcarg { $fplist = append($fplist, $funcarg.fp) }
       | {$fplist = make([]environment.FuncArg, 0)}
    ;


funcarg returns [environment.FuncArg fp]

: ID DOSPTOS expr {
                    $fp = environment.FuncArg{
                        Id:   $ID.text,
                        RealId: "",
                        Value: $expr.e,
                        Reference: false,
                    }
                  }

| expr {
                      $fp = environment.FuncArg{
                          Id:   "",
                          RealId: "",
                          Value: $expr.e,
                          Reference: false,
                      }
                    }

| ID DOSPTOS AMPERSAND expr {
                      $fp = environment.FuncArg{
                          Id:   $ID.text,
                          RealId: $expr.text,
                          Value: $expr.e,
                          Reference: true,
                      }
                    }

| AMPERSAND expr {
                        $fp = environment.FuncArg{
                            Id:   "",
                            RealId: $expr.text,
                            Value: $expr.e,
                            Reference: true,
                        }
                      }


;



// ------------------------------------------------------
funcparameterlist returns [[]environment.FuncParam fplist]
@init {
    $fplist = make([]environment.FuncParam, 0)
}
     : funcparameter { $fplist = append($fplist, $funcparameter.fp) }
       COMA a=funcparameterlist  { $fplist = append($fplist, $a.fplist...) }
       | funcparameter { $fplist = append($fplist, $funcparameter.fp) }
       | {$fplist = make([]environment.FuncParam, 0)}
    ;


funcparameter returns [environment.FuncParam fp]

: p=(ID|UNDERSCORE)  ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) {
                    $fp = environment.FuncParam{
                        Id:   $p.text,
                        SID: $ID.text,
                        Type: $typpe.text,
                        Reference: false,
                    }
                  }
// TODO ADD MATRIX
| p=(ID|UNDERSCORE) ID DOSPTOS  matrix_type  {
                        $fp = environment.FuncParam{
                            Id:   $p.text,
                            SID: $ID.text,
                            Type: $matrix_type.text,
                            Reference: false,
                        }
                      }




| p=(ID|UNDERSCORE) ID DOSPTOS RINOUT typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) {
                      $fp = environment.FuncParam{
                          Id:   $p.text,
                          SID: $ID.text,
                          Type: $typpe.text,
                          Reference: true,
                      }
                    }

| p=(ID|UNDERSCORE) ID DOSPTOS RINOUT matrix_type  {
                      $fp = environment.FuncParam{
                          Id:   $p.text,
                          SID: $ID.text,
                          Type: $matrix_type.text,
                          Reference: true,
                      }
                    }

//------------------------------------------

 | ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) {
                    $fp = environment.FuncParam{
                        Id:   "_",
                        SID: $ID.text,
                        Type: $typpe.text,
                        Reference: false,
                    }
                  }
|  ID DOSPTOS  matrix_type  {
                        $fp = environment.FuncParam{
                            Id:   "_",
                            SID: $ID.text,
                            Type: $matrix_type.text,
                            Reference: false,
                        }
                      }




|  ID DOSPTOS RINOUT typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER|ID) {
                      $fp = environment.FuncParam{
                          Id:   "_",
                          SID: $ID.text,
                          Type: $typpe.text,
                          Reference: true,
                      }
                    }

|  ID DOSPTOS RINOUT matrix_type  {
                      $fp = environment.FuncParam{
                          Id:   "_",
                          SID: $ID.text,
                          Type: $matrix_type.text,
                          Reference: true,
                      }
                    }


;





decmatrix returns [interfaces.Instruction newmatrix]
: RVAR ID IG manualdef  {$newmatrix = instructions.NewMatrixDec($RVAR.line, $RVAR.pos,$ID.text,nil,$manualdef.newmanualdef)}
| RVAR ID DOSPTOS matrix_type IG manualdef  {$newmatrix = instructions.NewMatrixDec($RVAR.line, $RVAR.pos,$ID.text,$matrix_type.text,$manualdef.newmanualdef)}
| RVAR ID IG repeatingvector  {$newmatrix = instructions.NewMatrixDec($RVAR.line, $RVAR.pos,$ID.text,nil,$repeatingvector.newrepeatingvec)}
| RVAR ID DOSPTOS matrix_type IG repeatingvector  {$newmatrix = instructions.NewMatrixDec($RVAR.line, $RVAR.pos,$ID.text,$matrix_type.text,$repeatingvector.newrepeatingvec)}
;


attrlist returns [[]string atrlist]
@init {
    $atrlist = []string{}
}
     : attr { $atrlist = append($atrlist, $attr.atr) }
       PTO a=attrlist  { $atrlist = append($atrlist, $a.atrlist...) }
       | attr { $atrlist = append($atrlist, $attr.atr) }
    ;


attr returns [string atr]
: ID {$atr=$ID.text}

;

// TODO maybe fix this?
structexp returns [interfaces.Expression structexxp ]
: ID LLAVEIZQ keyvaluelist LLAVEDER  {$structexxp = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $keyvaluelist.kvlist)}
;

keyvaluelist returns [[]environment.KeyValue kvlist]
@init {
    $kvlist = []environment.KeyValue{}
}
     : keyvalue { $kvlist = append($kvlist, $keyvalue.kv) }
       COMA a=keyvaluelist  { $kvlist = append($kvlist, $a.kvlist...) }
       | keyvalue { $kvlist = append($kvlist, $keyvalue.kv) }
       | {}
    ;

keyvalue returns [environment.KeyValue kv]

: ID DOSPTOS expr {
                    $kv = environment.KeyValue{
                        Key:   $ID.text,
                        Value: $expr.e,
                    }
                  }

;




callfuncexp returns [interfaces.Expression newcallfuncexp]
: ID PARIZQ funcarglist PARDER { $newcallfuncexp = expressions.NewCallFuncExp($ID.line, $ID.pos, $ID.text, $funcarglist.fplist) }

;

callfuncins returns [interfaces.Instruction newcallfuncins]
: ID PARIZQ funcarglist PARDER { $newcallfuncins = instructions.NewCallFuncInst($ID.line, $ID.pos, $ID.text, $funcarglist.fplist) }

;
// EXPRESSIONS -----------------------------------------------------------------

cast returns [interfaces.Expression newcast]
: RINT PARIZQ expr PARDER {$newcast = expressions.NewCast($RINT.line, $RINT.pos, "Int", $expr.e)}
| RSTRING PARIZQ expr PARDER {$newcast =expressions.NewCast( $RSTRING.line, $RSTRING.pos, "String", $expr.e)}
| RFLOAT PARIZQ expr PARDER {$newcast = expressions.NewCast($RFLOAT.line, $RFLOAT.pos, "Float", $expr.e)}
;

selfattributeaccess returns [interfaces.Expression newselfaccsess]
 : RSELF PTO ID {$newselfaccsess = expressions.NewSelfAccess($RSELF.line, $RSELF.pos, $ID.text)}
 ;


expr returns [interfaces.Expression e]
: cast  { $e = $cast.newcast }
|PARIZQ expr PARDER { $e = $expr.e }
| op=(SUB | NOT) left=expr { $e = expressions.NewUnaryOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text) }
| left=expr op=(MUL|DIV|MODULE) right=expr { $e = expressions.NewArithmeticOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }//TODO ACA DEBERIA IR EL MODULO
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewArithmeticOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR|MEN_IG|MENOR) right=expr { $e = expressions.NewRelationalOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewRelationalOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewBooleanOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewBooleanOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| vectorstructaccess {$e = $vectorstructaccess.vecstructaccess}
| callfuncexp {$e = $callfuncexp.newcallfuncexp}
| structexp {$e = $structexp.structexxp}
| structaccess {$e = $structaccess.saccess}
| isemptyvec {$e = $isemptyvec.newisemptyvec}
| countvec {$e = $countvec.newcountvec}
| vectoraccess {$e = $vectoraccess.newvecaccess}
| selfattributeaccess {$e = $selfattributeaccess.newselfaccsess}
| OBRA arguments CBRA {$e = expressions.NewVector($OBRA.line, $OBRA.pos, $arguments.args) }
| ID                        { $e = expressions.NewVariableAccess($ID.text) }
| NUMBER
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        var chari = str[1:len(str)-1]
        chari = strings.ReplaceAll(chari, "\\n","\n")
        chari = strings.ReplaceAll(chari, "\\r","\r")
        chari = strings.ReplaceAll(chari, "\\t","\t")
        chari = strings.ReplaceAll(chari, "\\\\","\\")
        chari = strings.ReplaceAll(chari, "\\\"","\"")
        if (len(str) == 3){

            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, chari,environment.CHAR)
        } else{

            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, chari,environment.STRING)
        }

    }
| RTRUE { $e = expressions.NewPrimitive($RTRUE.line, $RTRUE.pos, true, environment.BOOLEAN) }
| RFALSE { $e = expressions.NewPrimitive($RFALSE.line, $RFALSE.pos, false, environment.BOOLEAN) }
| RNIL  { $e = expressions.NewPrimitive($RFALSE.line, $RFALSE.pos, nil, environment.NULL) }
;




// --------------- Tokens
// types
RINT: 'Int';
RFLOAT: 'Float';
RBOOL: 'Bool';
RSTRING: 'String';
RCHARACTER: 'Character';

// reserved words
RTRUE: 'true';
RFALSE: 'false';
RPRINT: 'print';
RIF: 'if';
RELSE: 'else';
RWHILE: 'while';
RVAR: 'var';
RLET: 'let';
RNIL: 'nil';
RBREAK: 'break';
RCONTINUE:'continue';
RAPPEND : 'append'  ;
RREMOVELAST : 'removeLast';
RRAT        : 'at';
RREMOVEAT   : 'remove';
RISEMPTY :'isEmpty';
RCOUNT : 'count';
RSWITCH :'switch';
RCASE : 'case';
RDEFAULT : 'default';
RFOR : 'for';
RIN : 'in';
RREPEATING: 'repeating';
RSTRUCT: 'struct';
RGUARD : 'guard';
RFUNC : 'func';
RRETURN : 'return';
RINOUT : 'inout';
RMUTATING: 'mutating';
RSELF : 'self' ;

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"' ('\\' ('\\' | '"' | 'r'| 't'| 'n') | ~('\\' | '"'))* '"';

ID: [a-zA-Z_][a-zA-Z0-9_]*;

// symbols
ARROW :         '->';
UNARYPLUS:      '+=';
UNARYMINUS:     '-=';
DIF:            '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:             '=';
MAY_IG:         '>=';
MEN_IG:         '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
OBRA:           '[';
CBRA:           ']';
PTOCOMA:        ';';
DOSPTOS:        ':';
MODULE:         '%' ;
COMA:           ',' ;
QM:             '?' ;
PTO:            '.';
AMPERSAND :     '&';
UNDERSCORE :    '_';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//fragment
//ESC_SEQ
//    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
//    ;
