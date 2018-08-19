# ApiGo (https://github.com/eGenGuru/ApiGo)
# -----------------------------------------
# script sample for Perl

use JSON;

my %msg = ('message' => 'Hello World!');
my $json = encode_json \%msg;
print $json